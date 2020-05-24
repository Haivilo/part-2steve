# Part2
## Instruction

First download Go from https://golang.org/dl/. Here is an instruction document: https://golang.org/doc/install.

Then download the project and change to its directory by:

```bash
$git clone https://github.com/Haivilo/part-2steve.git
$cd part-2steve
```

Run Go:
```bash
$go run *.go
```

You should be able to see the webpage on localhost:8000.

## Potential issues
When running Go, there might be issues such as "pattern matches no files: `*.html`". This might happen because m.go cannot read HTML files using relevant directories. This can be resolved if you change the terminal's directory to the folder that cloned using git. 

## Notes
The account name and password is in account.json

I used cookies for login to save user's information on their local web browser. However my method is very simple, and this might be a security risk. Usually we are supposed to generate tokens randomly and dynamically delete them at backend server, it's doable but it will take a long time to finish. 

When logged in, the server will automatically redirect to the upload page unless logout.

When switching to French, I just added an "FR" in front of original texts for illustrating purpose. I used key-value pairs in every HTML page to store English and French word that's corresponding to each label.
