* It should accept links to documents provided in the POST request and download them.
* It should process and index the downloaded requests.
* It should handle the search queries and respond with a list of documents with snippets containing the search term.
* The result order should be in the order of greater occurence of search terms in the document.


2 Main Components:
* Concierge: Responsible for indexing and returning the list of documents for search queries.
* Librarian: Responsible for handling user interaction.

* These 2 will run as separate servers.
* We will have 3 instances of librarian.
* docker run works well for single images but for a network it might get complicated. So we'll use docker-compose.yaml