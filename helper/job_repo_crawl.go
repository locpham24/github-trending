package helper

import (
	"context"
	"fmt"
	"github.com/locpham24/github-trending/c_errors"
	"github.com/locpham24/github-trending/model"
	"github.com/locpham24/github-trending/repository"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gocolly/colly"
	repo "github.com/locpham24/github-trending/repository"
)

func CrawlRepo(githubRepository repo.GithubRepo) {
	c := colly.NewCollector()

	repos := make([]model.Repo, 0, 30)
	// Find and visit all links
	c.OnHTML("article[class=Box-row]", func(e *colly.HTMLElement) {
		rawName := e.ChildText("h1.h3 > a")
		name := strings.Replace(rawName, "\n", "", -1)
		name = strings.Replace(name, " ", "", -1)
		fmt.Println("name: ", name)

		rawColor := e.ChildAttr(".f6 .repo-language-color", "style")
		re := regexp.MustCompile("#[a-zA-Z0-9_]+")
		match := re.FindStringSubmatch(rawColor)
		color := ""
		if len(match) > 0 {
			color = match[0]
		}
		fmt.Println("color: ", color)

		rawURL := e.ChildAttr("h1.h3 > a", "href")
		url := "https://github.com/" + rawURL
		fmt.Println("url: ", url)

		rawLang := e.ChildText(".f6 span[itemprop=programmingLanguage]")
		lang := rawLang
		fmt.Println("lang: ", lang)

		stars := ""
		fork := ""
		e.ForEach(".mt-2 a", func(index int, element *colly.HTMLElement) {
			if index == 0 {
				rawStars := element.Text
				stars = strings.TrimSpace(rawStars)
			} else if index == 1 {
				rawFork := element.Text
				fork = strings.TrimSpace(rawFork)
			}
		})

		fmt.Println("stars: ", stars)
		fmt.Println("fork: ", fork)

		rawStarsToday := e.ChildText(".mt-2 .float-sm-right")
		starsToday := strings.TrimSpace(rawStarsToday)
		fmt.Println("starsToday: ", starsToday)

		var rawAuthors []string
		e.ForEach(".mt-2 .mr-3 a", func(index int, element *colly.HTMLElement) {
			srcImage := element.ChildAttr("img", "src")
			rawAuthors = append(rawAuthors, srcImage)
		})
		authors := strings.Join(rawAuthors, ",")
		fmt.Println("authors: ", authors)

		githubRepo := model.Repo{
			Name:        name,
			Description: "",
			Url:         url,
			Color:       color,
			Lang:        lang,
			Fork:        fork,
			Stars:       stars,
			StarsToday:  starsToday,
			Author:      authors,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}
		repos = append(repos, githubRepo)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		workerPool := NewWorkerPool(runtime.NumCPU())

		for _, repo := range repos {
			workerPool.Submit(&RepoProcess{
				repo:       repo,
				githubRepo: githubRepository,
			})
		}
		workerPool.Close()
		workerPool.Start()
	})

	c.Visit("https://github.com/trending")
}

type RepoProcess struct {
	repo       model.Repo
	githubRepo repository.GithubRepo
}

func (r *RepoProcess) Process() {
	cacheRepo, err := r.githubRepo.SelectRepoByName(context.Background(), r.repo.Name)
	if err == c_errors.RepoNotFound {
		fmt.Println("Add: ", r.repo.Name)
		_, err = r.githubRepo.SaveRepo(context.Background(), r.repo)
		if err != nil {
		}
		return
	}

	if cacheRepo.Stars != r.repo.Stars ||
		cacheRepo.Fork != r.repo.Fork ||
		cacheRepo.StarsToday != r.repo.StarsToday {
		fmt.Println("Update: ", r.repo.Name)
		_, err = r.githubRepo.UpdateRepo(context.Background(), r.repo)
		if err != nil {
		}
	}

	return
}
