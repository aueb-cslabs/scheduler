# aueb-cslabs-scheduler
This piece of software was created for the purpose of generating a schedule for the AUEB CSLabs.

## Run the schedule creator
To properly run the scheduler, you need to:
* Have the GoLang library installed on your machine.
* Have `github.com/SebastiaanKlippert/go-wkhtmltopdf` installed, as well as [wkhtmltopdf](https://wkhtmltopdf.org/) itself.
    * You can install the library by running `go get -u github.com/SebastiaanKlippert/go-wkhtmltopdf`
* `GOPATH` needs to point the directory you have cloned/downloaded the repository.
* Have created the custom rules method `func(admin model.Admin, time model.DayTime, model.lab int) bool`.
    * The created method, if it exists, must be pointed at `model.CustomBlockRule`
    * You can modify lines 21 and 89 of [src\aueb.gr\cslabs\scheduler\scheduler.go](src\aueb.gr\cslabs\scheduler\scheduler.go)
    * If there is none, you can remove those lines completely with no problems.
* Have a preferences CSV file. You can create it using [this](template.xlsx) spreadsheet as a template.
* After all these steps are complete you can run `go build aueb.gr/cslabs/scheduler` to build it!

And boom, the `scheduler.exe` or `scheduler` binary file is now created and ready to be used.

Execute `scheduler --help` to see all available options.