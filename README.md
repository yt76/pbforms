# PBForms: Toy example to generate HTML form from a protocol buffer message.


    $ protoc myform.proto --go_out=forms/
    
    $ go build cmd/dump.go
    $ ./dump > form.html

    $ go build cmd/server.go
    $ ./server &
    $ firefox http://localhost:8080/


TODO:

* Populate initial values
* Field name and HTML text handling (and i18n)
* Required/optional value
* Flexible HTML styling
* Multiple choices
* Radio button
* Option menu
* Repeated field
* Hidden field (especially CSRF token)
* Other fancy features
