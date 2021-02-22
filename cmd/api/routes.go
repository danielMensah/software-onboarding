package main

func initializeRoutes() {
	router.GET("/item", GetItemsHandler)
	router.GET("/item/:id", GetItemByIdHandler)
}
