terraform {
  required_providers {
    twc = {
      source = "tf.timeweb.cloud/timeweb-cloud/timeweb-cloud"
    }
    random = {
      source = "hashicorp/random"
      version = "3.7.1"
    }
  }
  required_version = ">= 1.4.4"
}

provider "twc" {
    token="eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCIsImtpZCI6IjFrYnhacFJNQGJSI0tSbE1xS1lqIn0.eyJ1c2VyIjoiY2MxOTk1MCIsInR5cGUiOiJhcGlfa2V5IiwiYXBpX2tleV9pZCI6IjNjMzVhNjBkLWMxYzUtNDQ2NS05Zjc4LTNiZDE3NmIwZTMxOSIsImlhdCI6MTc0MDg0OTEyOH0.c3_pNvSQ1Fc90ME0hmfFBLxVuMs_FPjItilAuYvLEfqxBc93_b1mjADjEKCeMj3jfGZHwVkmQDGNgGFCnUx2BjKpncgpeHmHojWylkbe8Zvqw5NB6FJ_6c1uJQ5FrpoffykL6smpP639NvexdAcJEOGOVA4um06ofzhOe7TW_053TkNcdPh5lmmDoIzghBRYYlATIRYtzbKWys3V5df-SgAjMl3z0LMPbFN2VVNazGJO0MSj2jRA4rdUhrJDExyQrdIE7ADbDSlOUxFCODwXsdqNoT9IopfJmFOAGLzM7nkk3nejIKNb3nkc5eFL0wLZrBCX1rKvlxmZh0kLSLtjMm_ZX798TqTlTVNhtghQMQeoFOXzG4VofW8lLqIx7v4iiAnwWdjhVvkvdAX-87Gar9Ad6ImZkCzFiVZkK67S4S1QGb3V7X4NIebV2UJPJl5bMxCZoaXvzahcsF3c53EkGDbB03HXOqU2Yck17tBd4SykhMP-bpUEFKcI9EBKs2lS"
}

provider "random" {
}