# clean-architecture-golang

Learning Uncle Bob’s concept of The Clean Architecture.

## Description

Creation of a simple yet complete Go application,  applied the concepts of The Clean Architecture.

The application is a very (very!) rudimentary but working online shopping system which allows to retrieve a list of items, belonging to an order, through an HTTP web service.

If there is duplicate code, no error handling and many other smells, it is because this example is neither code style nor design patterns: it is an example of the architecture of application, and therefore I took the liberty to create very simple that just needs to be simple and understandable, not smart and smart - oh and yes, I'm still a beginner Go, which proves it.



* Domain:
    * Customer entity
    * Item entity
    * Order entity

* Use Cases:
    * User entity
    * Use case: Add Item to Order
    * Use case: Get Items in Order
    * Use case: Admin adds Item to Order

* Infrastructure:
    * Web Services for Item/Order handling
    * Repositories for Use Cases and Domain entities persistence
    * The Database
    * Code that handles DB connections
    * The HTTP server
    * Go Standard Library