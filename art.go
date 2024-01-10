package ghart

import (
	"os"
	"os/exec"
	"time"
)

type Point struct {
	X int
	Y int
}

type Config struct {
	Name string
	Email string
}

func GenerateArt(art []Point, config Config, path string) error {
	dates := getDates(art)

	if err := initRepo(config, path); err != nil {
		return err
	}

	for _, date := range dates {
		if err := commit(date); err != nil {
			return err
		}
	}

	return nil
}

func getDates(points []Point) []time.Time {
	dates := []time.Time{}

	for _, p := range points {
		dates = append(dates, pointToDate(p))
	}

	return dates
}

func pointToDate(p Point) time.Time {
	distance := ((p.X * 7) + 1) + p.Y
	// there are no pixels before this pixel
	firstDay := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.Local).Weekday()
	dayOfYear := distance - int(firstDay)
	return time.Date(2019, time.January, dayOfYear, 0, 0, 0, 0, time.Local)
}

func initRepo(config Config, path string) error {
	if err := os.Chdir(path); err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := configureGit(config); err != nil {
		return err
	}

	if err := os.WriteFile("README.md", []byte("gh-art"), 0666); err != nil {
		return err
	}

	cmd = exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func configureGit(config Config) error {
	cmd := exec.Command("git", "config", "--local", "user.name", config.Name)
	if err := cmd.Run(); err != nil {
		return err
	}

	cmd = exec.Command("git", "config", "--local", "user.email", config.Email)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// TODO: make sure we cd or just give path
func commit(date time.Time) error {
	cmd := exec.Command("git", "commit", "-m", "gh-art", "--allow-empty", "--date", date.String())
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
