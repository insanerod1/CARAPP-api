#CARAPP-api (still working on a better name)
###The api used in this godot game https://github.com/insanerod1/3390-Project-3

This is a quick web api that uses Go/Gin as the framework to route http requests. In the godot game it's used to randomize death comments depending on the obstacle that the player died too. It is also use to generate the comments of the driving assistant in the game based on what obstacle is incoming.

##HOW TO INSTALL
***Note these instructions are valid for go version 1.27.3 and Gin version 1.12.0 which is the latest release at the time of writing this***
If you are on ***Windows*** or ***Mac*** you can use this link to download Golang to your device https://go.dev/

If you are on Linux (Ubuntu/Debian) you can run this command in the terminal

> sudo apt install golang

After installing clone the repo and open the project folder, then in the terminal type

> go run .

It should start the api you will know everything is running well if it says
> [GIN-debug] Listening and serving HTTP on localhost:8080

##HOW TO USE

To make get requests simply run the api on the localhost:8080 and in a client or browser use

> http://localhost:8080/comments/

If you get **404 page not found** that's a good thing since we haven't added anything else to the header

##For death comments use 

>./comments/death?obstacle=

the available values for obstacle are
-car
-truck
-truck_blockade
-tire
-flaming_tire
-cone
-cop

##For assistant comments use

>./comments/assistant?event=

the available values for event are
-idle
-speeding
-truck_blockade
-tires
-cone
-cop

