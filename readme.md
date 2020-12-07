# **DM-RECIPE**

tech challenge received from [Delivery Much](http://fordevs.herokuapp.com/api-docs) using the principles of [Clean Architecture](https://www.google.com/search?q=clean+architecture+book&rlz=1C1CHBD_pt-BRBR926BR926&sxsrf=ALeKk01iHS5sS9zlNO8Qsn8Dsh6Xo4ddJA:1607287301707&tbm=isch&source=iu&ictx=1&fir=yJSXMdiMorUknM%252CSUw2kUqGpluSAM%252C%252Fg%252F11c1pmnj82&vet=1&usg=AI4_-kS7Pukfn5YKGPdL0PLU1-Ed8fG-1Q&sa=X&ved=2ahUKEwiZ6K3lm7rtAhWTHbkGHXiXApIQ_B16BAgoEAI&biw=1920&bih=937#imgrc=yJSXMdiMorUknM)

  

## How to start

You'll need docker-compose to run like this

```
docker-compose up
```

or you can build it by your-self

```
docker build -t recipe .
```

## API
Has only one endpoint, which must respect the following call:

`
http://{HOST}/recipes/?i={ingredient_1},{ingredient_2},{ingredient_3}
`

Example:

`
http://127.0.0.1:8080/recipes/?i=onion,tomato
`

Return:

```
{
	"keywords": ["onion", "tomato"],
	"recipes": [{
		"title": "Greek Omelet with Feta",
		"ingredients": ["eggs", "feta cheese", "garlic", "red onions", "spinach", "tomato", "water"],
		"link": "http://www.kraftfoods.com/kf/recipes/greek-omelet-feta-104508.aspx",
		"gif": "https://media.giphy.com/media/xBRhcST67lI2c/giphy.gif"
	   },{
		"title": "Guacamole Dip Recipe",
		"ingredients": ["avocado", "onions", "tomato"],
		"link":"http://cookeatshare.com/recipes/guacamole-dip-2783",
		"gif":"https://media.giphy.com/media/I3eVhMpz8hns4/giphy.gif"
	   }
	]
}
```
