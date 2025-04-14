$exeDownload = "https://github.com/wlfstn/lsf/releases/download/V1.3/lsf.exe"
$destinationDir = "$env:USERPROFILE\wlfstn\lsf"
$destinationFile = Join-Path $destinationDir "lsf.exe"

if (-not (Test-Path $destinationDir)) {
	New-Item -ItemType Directory -Path $destinationDir -Force | Out-Null
}

Invoke-WebRequest -Uri $exeDownload -OutFile $destinationFile

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")
if (-not $userPath) {
	$userPath = ""
}

if ($userPath -split ";" -contains $destinationDir) {
	Write-Host "lsf directory is already in PATH."
} else {
	$addToPath = Read-Host "Would you like to add lsf to your PATH? (y/n)"
	if ($addToPath -match '^[Yy]$') {
		$newPath = if ($userPath -eq "") {
			$destinationDir
		} else {
			"$userPath;$destinationDir"
		}
		[Environment]::SetEnvironmentVariable("Path", $newPath, "User")
		Write-Host "Added $destinationDir to user PATH."
	} else {
		Write-Host "Installed, but skipped adding to PATH."
	}
}