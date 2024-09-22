package database

import (
	"fmt"
	"time"
)

var SampleAnime = []Anime{
	{
		ID:          "bfa7bcb0-6c97-4423-a1bc-92f64cfaab7f",
		Title:       "Attack on Titan",
		ReleaseDate: parseDate("2013-04-07"),
		Rating:      9.0,
		Genre:       []string{"Action", "Drama", "Fantasy"},
		Type:        TV,
		Episodes:    87,
		Description: "\tEren Yeager once lived a peaceful life behind the towering walls that protected humanity from the terrifying Titans. But that peace was shattered when a massive Titan broke through the outer wall, destroying his home and taking his mother’s life. Determined to rid the world of these monsters, Eren joins the military alongside his friends to fight back.\n\tHowever, as the battles intensify and Eren uncovers hidden secrets about the Titans and the world, he begins to question who the real enemy is. With humanity’s survival hanging in the balance, Eren must confront betrayal, conspiracy, and his own inner demons as he fights to uncover the truth.",
		Studio:      []string{"Wit Studio", "MAPPA"},
		Duration:    24,
		Status:      Finished,
		ESRB:        M,
	},
	{
		ID:          "49c02c34-2300-4f15-aea3-1a9475f0b4a9",
		Title:       "Your Name",
		ReleaseDate: parseDate("2016-08-26"),
		Rating:      8.9,
		Genre:       []string{"Romance", "Supernatural", "Drama"},
		Type:        Movie,
		Episodes:    1,
		Description: "Two teenagers find themselves swapping bodies and must find a way to meet each other.",
		Studio:      []string{"CoMix Wave Films"},
		Duration:    106,
		Status:      Finished,
		ESRB:        T,
	},
	{
		ID:          "b2b0a62f-9fcb-4fa8-a07b-8b0f6b27e56a",
		Title:       "My Hero Academia",
		ReleaseDate: parseDate("2016-04-03"),
		Rating:      8.0,
		Genre:       []string{"Action", "Comedy", "Superhero"},
		Type:        TV,
		Episodes:    138,
		Description: "In a world where nearly everyone has superpowers, one boy without powers strives to become a hero.",
		Studio:      []string{"Bones"},
		Duration:    24,
		Status:      Airing,
		ESRB:        T,
	},
	{
		ID:          "d8a1f6f1-6e2b-42be-8c56-0edb20d3f19d",
		Title:       "Spirited Away",
		ReleaseDate: parseDate("2001-07-20"),
		Rating:      8.6,
		Genre:       []string{"Adventure", "Fantasy", "Supernatural"},
		Type:        Movie,
		Episodes:    1,
		Description: "A young girl becomes trapped in a strange world of spirits and must find a way to escape.",
		Studio:      []string{"Studio Ghibli"},
		Duration:    125,
		Status:      Finished,
		ESRB:        T,
	},
	{
		ID:          "be215d92-1a8f-49e4-b926-58c1e6db7300",
		Title:       "One Piece",
		ReleaseDate: parseDate("1999-10-20"),
		Rating:      8.7,
		Genre:       []string{"Action", "Adventure", "Fantasy"},
		Type:        TV,
		Episodes:    1071,
		Description: "Monkey D. Luffy and his pirate crew search for the legendary One Piece treasure.",
		Studio:      []string{"Toei Animation"},
		Duration:    24,
		Status:      Airing,
		ESRB:        T,
	},
	{
		ID:          "9d8c3c1e-3f9d-4375-ae43-8dbe658f37ad",
		Title:       "Demon Slayer",
		ReleaseDate: parseDate("2019-04-06"),
		Rating:      8.7,
		Genre:       []string{"Action", "Supernatural", "Historical"},
		Type:        TV,
		Episodes:    44,
		Description: "A young boy becomes a demon slayer to avenge his family and save his sister.",
		Studio:      []string{"ufotable"},
		Duration:    24,
		Status:      Airing,
		ESRB:        M,
	},
	{
		ID:          "a6a755f2-0545-41be-92d8-92d3f056b39a",
		Title:       "Cowboy Bebop",
		ReleaseDate: parseDate("1998-04-03"),
		Rating:      8.9,
		Genre:       []string{"Action", "Sci-Fi", "Space"},
		Type:        TV,
		Episodes:    26,
		Description: "A group of bounty hunters travel through space, catching criminals and facing their pasts.",
		Studio:      []string{"Sunrise"},
		Duration:    24,
		Status:      Finished,
		ESRB:        M,
	},
}

// Utility function to parse the release date
func parseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
	}
	return date
}
