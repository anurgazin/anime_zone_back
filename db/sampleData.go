package database

// import (
// 	"fmt"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var SampleAnime = []Anime{
// 	{
// 		ID:          parseId("bfa7bcb06c974423a1bc92f6"),
// 		Title:       "Attack on Titan",
// 		ReleaseDate: parseDate("2013-04-07"),
// 		Rating:      9.0,
// 		Genre:       []string{"Action", "Drama", "Fantasy"},
// 		Type:        TV,
// 		Episodes:    87,
// 		Description: "\tEren Yeager once lived a peaceful life behind the towering walls that protected humanity from the terrifying Titans. But that peace was shattered when a massive Titan broke through the outer wall, destroying his home and taking his mother’s life. Determined to rid the world of these monsters, Eren joins the military alongside his friends to fight back.\n\tHowever, as the battles intensify and Eren uncovers hidden secrets about the Titans and the world, he begins to question who the real enemy is. With humanity’s survival hanging in the balance, Eren must confront betrayal, conspiracy, and his own inner demons as he fights to uncover the truth.",
// 		Studio:      []string{"Wit Studio", "MAPPA"},
// 		Duration:    24,
// 		Status:      Finished,
// 		ESRB:        M,
// 	},
// 	{
// 		ID:          parseId("49c02c3423004f15aea31a94"),
// 		Title:       "Your Name",
// 		ReleaseDate: parseDate("2016-08-26"),
// 		Rating:      8.9,
// 		Genre:       []string{"Romance", "Supernatural", "Drama"},
// 		Type:        Movie,
// 		Episodes:    1,
// 		Description: "Two teenagers find themselves swapping bodies and must find a way to meet each other.",
// 		Studio:      []string{"CoMix Wave Films"},
// 		Duration:    106,
// 		Status:      Finished,
// 		ESRB:        T,
// 	},
// 	{
// 		ID:          parseId("b2b0a62f9fcb4fa8a07b8b0f"),
// 		Title:       "My Hero Academia",
// 		ReleaseDate: parseDate("2016-04-03"),
// 		Rating:      8.0,
// 		Genre:       []string{"Action", "Comedy", "Superhero"},
// 		Type:        TV,
// 		Episodes:    138,
// 		Description: "In a world where nearly everyone has superpowers, one boy without powers strives to become a hero.",
// 		Studio:      []string{"Bones"},
// 		Duration:    24,
// 		Status:      Airing,
// 		ESRB:        T,
// 	},
// 	{
// 		ID:          parseId("d8a1f6f16e2b42be8c560edb"),
// 		Title:       "Spirited Away",
// 		ReleaseDate: parseDate("2001-07-20"),
// 		Rating:      8.6,
// 		Genre:       []string{"Adventure", "Fantasy", "Supernatural"},
// 		Type:        Movie,
// 		Episodes:    1,
// 		Description: "A young girl becomes trapped in a strange world of spirits and must find a way to escape.",
// 		Studio:      []string{"Studio Ghibli"},
// 		Duration:    125,
// 		Status:      Finished,
// 		ESRB:        T,
// 	},
// 	{
// 		ID:          parseId("be215d921a8f49e4b92658c1"),
// 		Title:       "One Piece",
// 		ReleaseDate: parseDate("1999-10-20"),
// 		Rating:      8.7,
// 		Genre:       []string{"Action", "Adventure", "Fantasy"},
// 		Type:        TV,
// 		Episodes:    1071,
// 		Description: "Monkey D. Luffy and his pirate crew search for the legendary One Piece treasure.",
// 		Studio:      []string{"Toei Animation"},
// 		Duration:    24,
// 		Status:      Airing,
// 		ESRB:        T,
// 	},
// 	{
// 		ID:          parseId("9d8c3c1e3f9d4375ae438dbe"),
// 		Title:       "Demon Slayer",
// 		ReleaseDate: parseDate("2019-04-06"),
// 		Rating:      8.7,
// 		Genre:       []string{"Action", "Supernatural", "Historical"},
// 		Type:        TV,
// 		Episodes:    44,
// 		Description: "A young boy becomes a demon slayer to avenge his family and save his sister.",
// 		Studio:      []string{"ufotable"},
// 		Duration:    24,
// 		Status:      Airing,
// 		ESRB:        M,
// 	},
// 	{
// 		ID:          parseId("a6a755f2054541be92d892d3"),
// 		Title:       "Cowboy Bebop",
// 		ReleaseDate: parseDate("1998-04-03"),
// 		Rating:      8.9,
// 		Genre:       []string{"Action", "Sci-Fi", "Space"},
// 		Type:        TV,
// 		Episodes:    26,
// 		Description: "A group of bounty hunters travel through space, catching criminals and facing their pasts.",
// 		Studio:      []string{"Sunrise"},
// 		Duration:    24,
// 		Status:      Finished,
// 		ESRB:        M,
// 	},
// }

// // Utility function to parse the release date
// func parseId(id string) primitive.ObjectID {
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return objID
// }
// func parseDate(dateStr string) time.Time {
// 	date, err := time.Parse("2006-01-02", dateStr)
// 	if err != nil {
// 		fmt.Printf("Error parsing date: %v\n", err)
// 	}
// 	return date
// }

// var SampleCharacters = []Character{
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Eren",
// 		LastName:  "Yeager",
// 		Age:       19,
// 		FromAnime: []primitive.ObjectID{parseId("bfa7bcb06c974423a1bc92f6")},
// 		Gender:    "Male",
// 		Bio:       "Eren Yeager is a passionate and determined individual who joins the Scout Regiment to eliminate the Titans after witnessing the death of his mother. As the series progresses, Eren uncovers secrets about the Titans and himself that change his perspective on the world.",
// 		Status:    "dead",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Mitsuha",
// 		LastName:  "Miyamizu",
// 		Age:       17,
// 		FromAnime: []primitive.ObjectID{parseId("49c02c3423004f15aea31a94")},
// 		Gender:    "Female",
// 		Bio:       "Mitsuha Miyamizu is a high school girl from a rural town who longs for life in the bustling city of Tokyo. Her life changes drastically when she starts swapping bodies with a boy named Taki from Tokyo, leading to a mysterious and emotional journey.",
// 		Status:    "alive",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Izuku",
// 		LastName:  "Midoriya",
// 		Age:       16,
// 		FromAnime: []primitive.ObjectID{parseId("b2b0a62f9fcb4fa8a07b8b0f")},
// 		Gender:    "Male",
// 		Bio:       "Izuku Midoriya, also known as Deku, was born without a Quirk in a world where almost everyone has superpowers. Despite this, his determination to become a hero like his idol All Might leads him to inherit the powerful Quirk, One for All.",
// 		Status:    "alive",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Chihiro",
// 		LastName:  "Ogino",
// 		Age:       10,
// 		FromAnime: []primitive.ObjectID{parseId("d8a1f6f16e2b42be8c560edb")},
// 		Gender:    "Female",
// 		Bio:       "Chihiro Ogino is a young girl who gets trapped in the spirit world after her parents are transformed into pigs. She takes on the name Sen and works in a bathhouse for spirits to survive and find a way to rescue her parents.",
// 		Status:    "alive",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Monkey D.",
// 		LastName:  "Luffy",
// 		Age:       19,
// 		FromAnime: []primitive.ObjectID{parseId("be215d921a8f49e4b92658c1")},
// 		Gender:    "Male",
// 		Bio:       "Monkey D. Luffy is a pirate with the ability to stretch his body like rubber after consuming the Gomu Gomu no Mi devil fruit. His dream is to become the Pirate King by finding the legendary treasure known as One Piece.",
// 		Status:    "alive",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Tanjiro",
// 		LastName:  "Kamado",
// 		Age:       15,
// 		FromAnime: []primitive.ObjectID{parseId("9d8c3c1e3f9d4375ae438dbe")},
// 		Gender:    "Male",
// 		Bio:       "Tanjiro Kamado is a kind-hearted boy who becomes a demon slayer after his family is slaughtered by demons, and his sister Nezuko is turned into one. Tanjiro seeks to avenge his family and find a way to turn Nezuko back into a human.",
// 		Status:    "alive",
// 	},
// 	{
// 		ID:        primitive.NewObjectID(),
// 		FirstName: "Spike",
// 		LastName:  "Spiegel",
// 		Age:       27,
// 		FromAnime: []primitive.ObjectID{parseId("a6a755f2054541be92d892d3")},
// 		Gender:    "Male",
// 		Bio:       "Spike Spiegel is a laid-back bounty hunter with a mysterious past. He is a skilled fighter with a tragic love story that haunts him throughout his adventures aboard the spaceship Bebop.",
// 		Status:    "unknown",
// 	},
// }
