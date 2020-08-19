package server

import (
	"fmt"
	"github.com/SebastienBoisard/newsletters/internal/server/models"
	"github.com/jmoiron/sqlx"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func InitDatabase() (*sqlx.DB, error) {

	dbDriver := "mysql"
	dbUsername := viper.GetString("database.user.name")

	dbHost := viper.GetString("database.config.host")
	dbPort := viper.GetInt("database.config.port")

	dbPassword := viper.GetString("database.user.password")
	dbName := viper.GetString("database.name")

	db, err := sqlx.Connect(
		dbDriver,
		fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}

func CreateDatabase(db *sqlx.DB) {

	db.MustExec(models.CreateTableNewsletter)

	db.MustExec(models.CreateTableSubscriber)

	db.MustExec(models.CreateTableEpisode)

	db.MustExec(models.CreateTableBlock)

	db.MustExec(models.CreateTableHeader)

	db.MustExec(models.CreateTableFooter)

	db.MustExec(models.CreateTableSubscription)
}

func CleanDatabase(db *sqlx.DB) {

	db.MustExec(models.DeleteTableNewsletter)

	db.MustExec(models.DeleteTableSubscriber)

	db.MustExec(models.DeleteTableEpisode)

	db.MustExec(models.DeleteTableBlock)

	db.MustExec(models.DeleteTableHeader)

	db.MustExec(models.DeleteTableFooter)

	db.MustExec(models.DeleteTableSubscription)
}

func FillDatabase(db *sqlx.DB) {

	subscriber := addSubscriber(db, "john.smith@yopmail.com")

	newsletter := addNewsletter(db, "newsletter-01")

	addSubscription(db, newsletter, subscriber, "first-newsletter", "nKey", "sKey", -1, -1)

	header := addHeader(db,
		`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">

  <title>Oui la lettre #269 : Rien n'est plus patient que la nature.</title>
  <meta name="description" content="First Newsletter">
  <meta name="author" content="Me">

  <link href="https://unpkg.com/tailwindcss@^1.0/dist/tailwind.min.css" rel="stylesheet">
</head>
<body>
<div class="container mx-auto pt-4">`)

	footer := addFooter(db, `</div></body></html>`)

	episode := addEpisode(db, newsletter, 1, "2020-07-10", header, footer)

	addBlock(db, newsletter, episode, 1,
		`<div class="text-white bg-indigo-800 text-center py-2">
   <p class="text-xs">Pour vous désinscrire, cliquez ici</p>
   <p class="text-5xl pt-4">OUI! LA LETTRE</p>
   <p class="text-sm">DU 28 MAI AU 3 JUIN 2020</p>
</div>`)

	addBlock(db, newsletter, episode, 2,
		`<img alt="Refettorio" width="100%" style="width:100%;height:auto" src="/static/images/unnamed.jpg">
`)

	addBlock(db, newsletter, episode, 3,
		`	<div class="text-2xl my-8 flex justify-center">

	<p class="w-1/2">
		Toc toc toc, c’est la factrice. Envie d’accélérer ou de ralentir ? Voici quelques idées pour faire fructifier le temps retrouvé. Enterrée dans le jardin, la terre cuite diffuse l’eau nécessaire à la vie, longuement et doucement. Quand les aromates débordent du panier, ayons les gestes malins pour capter durablement leurs vertus. Vous trépignez de retrouver terrasses et tablées ? Les restaurateurs comptent dès maintenant sur notre soutien, pour rendre l’avenir léger.
	</p>

	</div>
`)
	addBlock(db, newsletter, episode, 4,
		`<div class="text-2xl my-8 flex justify-center">
	<img class="w-auto" src="/static/images/unnamed(1).jpg">
	</div>`)
	addBlock(db, newsletter, episode, 5,
		`<div class="my-8 flex  flex-col">
	<p class="self-center w-1/2 font-bold text-base">
		FONTAINE, JE BOIRAI DE TA MENTHE À L'EAU
	<a href="#" class="text-orange-600 underline">TRUCS ET ASTUCES</a>
	</p>
	<p class="self-center w-1/2 text-3xl font-bold">Menthra de saison</p>
	<p class="self-center w-1/2 text-2xl ">
		Soyons cigale et profitons maintenant de l’abondance de la menthe au jardin. Mais tout aussi fourmi pour ne pas se retrouver dépourvu lorsque la bise sera venue. C’est parti pour six idées parfumées et prévoyantes.
	</p>
	<div class=" w-1/2 self-center mt-4">
	<button class="text-base bg-orange-600 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded">Feuilleter</button>
	</div>
	</div>`)
	addBlock(db, newsletter, episode, 6,
		`	<div class="text-2xl my-8 flex justify-center">
	<img class="w-auto" src="/static/images/unnamed(2).jpg">
	</div>`)
	addBlock(db, newsletter, episode, 7,
		`<div class="my-8 flex  flex-col">
	<p class="self-center w-1/2 font-bold text-base">
		DES JARRES AU JARDIN
	<a href="#" class="text-orange-600 underline">IDÉES</a>
	</p>
	<p class="self-center w-1/2 text-3xl font-bold">De l’eau en pot, trésor à enterrer</p>
	<p class="self-center w-1/2 text-2xl ">
		Face aux sécheresses de plus en plus longues et régulières, jardiniers et amis des plantes cherchent des solutions durables et résilientes. Focus sur une technique ancestrale : l’irrigation par jarre.
	</p>
	<div class=" w-1/2 self-center mt-4">
	<button class="text-base bg-orange-600 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded">Se mouiller</button>
	</div>
	</div>`)
	addBlock(db, newsletter, episode, 8,
		`<div class="text-2xl my-8 flex justify-center">
	<img class="w-auto" src="/static/images/unnamed(3).jpg">
	</div>`)
	addBlock(db, newsletter, episode, 9,
		`<div class="my-8 flex  flex-col">
	<p class="self-center w-1/2 font-bold text-base">
		TOUS AU RESTAU
	<a href="#" class="text-orange-600 underline">GRAND ANGLE - NOUVELLE AQUITAINE</a>
	</p>
	<p class="self-center w-1/2 text-3xl font-bold">Rezto : des cartes cadeaux pour se serrer les coudes à table</p>
	<p class="self-center w-1/2 text-2xl ">
		Au Pays basque, la jeune start-up Rezto propose des cartes cadeaux solidaires à destination des restaurateurs. Un petit geste pour soutenir ces commerçants fortement impactés par la crise du Covid-19.
	</p>
	<div class=" w-1/2 self-center mt-4">
	<button class="text-base bg-orange-600 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded">Soutenir
	</button>
	</div>
	</div>
`)
	addBlock(db, newsletter, episode, 10,
		`<div class="text-2xl my-8 flex justify-center">
	<img class="w-auto" src="/static/images/unnamed(4).png">
	</div>
`)
	addBlock(db, newsletter, episode, 11,
		`	<div class="my-8 flex  flex-col">
	<p class="self-center w-1/2 font-bold text-base">
		CRESSON, FEUILLES POÉTIQUES
	<a href="#" class="text-orange-600 underline">RECETTE</a>
	</p>
	<p class="self-center w-1/2 text-2xl ">
		Rien n’est plus vert que le vert. Sur ce sujet haut en couleur, plus aucun doute ne sera permis après cette recette.
	</p>
	<div class=" w-1/2 self-center mt-4">
	<button class="text-base bg-orange-600 hover:bg-orange-600 text-white font-bold py-2 px-4 rounded">Souper
	</button>
	</div>
	</div>
`)
	addBlock(db, newsletter, episode, 12,
		`	<div class="text-white bg-indigo-800 justify-center py-2 flex ">
	<img class="px-4" src="/static/images/unnamed(10).png">
	<img class="px-4" src="/static/images/unnamed(11).png">
	<img class="px-4" src="/static/images/unnamed(12).png">
	<img class="px-4" src="/static/images/unnamed(13).png">
	</div>
`)
	addBlock(db, newsletter, episode, 13,
		`	<div class="my-4 flex flex-col items-center">
	<p>
		ENVIE DE NOUS PROPOSER UN SUJET D’ARTICLE ?
	</p>
	<p>
		N’hésitez pas à nous répondre, à commenter les articles et à nous
	proposer des idées de sujets en nous écrivant à ouimag@lrqdo.fr.
	</p>
	<p>
		Si vous souhaitez vous désinscrire de cette newsletter, cliquez ici
	</p>
	<div>
	<img src="/static/images/unnamed(14).png">
	</div>
	</div>
`)

}

func addSubscriber(db *sqlx.DB, email string) *models.Subscriber {

	subscriber := models.Subscriber{
		SubscriberID: ksuid.New().String(),
		Email:        email,
	}
	_, err := db.NamedExec(`
INSERT INTO Subscriber 
       (SubscriberID, Email) 
VALUES (:SubscriberID, :Email)
`,
		&subscriber)

	if err != nil {
		panic(err)
	}

	return &subscriber
}

func addNewsletter(db *sqlx.DB, newsletterName string) *models.Newsletter {

	newsletter := models.Newsletter{
		NewsletterID: ksuid.New().String(),
		Name:         newsletterName,
	}
	_, err := db.NamedExec(`
INSERT INTO Newsletter 
       (NewsletterID, Name) 
VALUES (:NewsletterID, :Name)
`,
		&newsletter)

	if err != nil {
		panic(err)
	}

	return &newsletter
}

func addHeader(db *sqlx.DB, content string) *models.Header {

	header := models.Header{
		HeaderID: ksuid.New().String(),
		Content:  content,
	}
	_, err := db.NamedExec(`
INSERT INTO Header 
       (HeaderID, Content) 
VALUES (:HeaderID, :Content)
`,
		&header)

	if err != nil {
		panic(err)
	}

	return &header
}

func addFooter(db *sqlx.DB, content string) *models.Footer {

	footer := models.Footer{
		FooterID: ksuid.New().String(),
		Content:  content,
	}
	_, err := db.NamedExec(`
INSERT INTO Footer 
       (FooterID, Content) 
VALUES (:FooterID, :Content)
`,
		&footer)

	if err != nil {
		panic(err)
	}

	return &footer
}

func addEpisode(db *sqlx.DB, newsletter *models.Newsletter, episodeID int, timestamp string, header *models.Header, footer *models.Footer) *models.Episode {

	creationDate, err := time.Parse("2006-01-02", timestamp)
	if err != nil {
		panic(err)
	}

	episode := models.Episode{
		EpisodeID:    episodeID,
		NewsletterID: newsletter.NewsletterID,
		CreationDate: creationDate,
		HeaderID:     header.HeaderID,
		FooterID:     footer.FooterID,
	}

	_, err = db.NamedExec(`
INSERT INTO Episode 
       (EpisodeID, NewsletterID, CreationDate, HeaderID, FooterID) 
VALUES (:EpisodeID, :NewsletterID, :CreationDate, :HeaderID, :FooterID)
`,
		&episode)

	if err != nil {
		panic(err)
	}

	return &episode
}

func addBlock(db *sqlx.DB, newsletter *models.Newsletter, episode *models.Episode, blockID int, content string) *models.Block {

	block := models.Block{
		NewsletterID: newsletter.NewsletterID,
		EpisodeID:    episode.EpisodeID,
		BlockID:      blockID,
		Content:      content,
	}
	_, err := db.NamedExec(
		`INSERT INTO Block 
       (NewsletterID, EpisodeID, BlockID, Content) 
VALUES (:NewsletterID, :EpisodeID, :BlockID, :Content)`,
		&block)

	if err != nil {
		panic(err)
	}

	return &block
}

func addSubscription(db *sqlx.DB, newsletter *models.Newsletter, subscriber *models.Subscriber, newsletterShortname string, newsletterKey string, subscriberKey string, startEpisodeID int, endEpisodeID int) *models.Subscription {

	subscription := models.Subscription{
		NewsletterID:        newsletter.NewsletterID,
		SubscriberID:        subscriber.SubscriberID,
		NewsletterShortname: newsletterShortname,
		NewsletterKey:       newsletterKey,
		SubscriberKey:       subscriberKey,
		StartEpisodeID:      startEpisodeID,
		EndEpisodeID:        endEpisodeID,
	}
	_, err := db.NamedExec(
		`INSERT INTO Subscription 
       (NewsletterID, SubscriberID, NewsletterShortname, NewsletterKey, SubscriberKey, StartEpisodeID, EndEpisodeID) 
VALUES (:NewsletterID, :SubscriberID, :NewsletterShortname, :NewsletterKey, :SubscriberKey, :StartEpisodeID, :EndEpisodeID)`,
		&subscription)

	if err != nil {
		panic(err)
	}

	return &subscription
}
