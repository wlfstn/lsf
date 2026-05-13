param (
	[string]$flag
)

switch ($flag) {
	"--debug" {
		Write-Host "Building..."
		go run .
	}
	"--build" {
		Write-Host "Building (debug)..."
		go build -o "./build/lsf.exe"
	}
	"--run" {
		Write-Host "Running..."
		& "./build/lsf.exe"
	}
	default {
		Write-Host "Usage: .\build.ps1 --build | --debug | --run"
	}
}
