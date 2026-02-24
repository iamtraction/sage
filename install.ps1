$ErrorActionPreference = "Stop"

$Arch = switch ($env:PROCESSOR_ARCHITECTURE) {
  "AMD64" { "amd64" }
  "ARM64" { "arm64" }
  default { Write-Error "Unsupported arch: $env:PROCESSOR_ARCHITECTURE" }
}

# resolve latest release version from GitHub
$Latest = Invoke-RestMethod -Uri "https://api.github.com/repos/iamtraction/sage/releases/latest"
$Tag = $Latest.tag_name
$Version = $Tag -replace '^v', ''
$Artifact = "sage_${Version}_windows_$Arch.zip"
$Url = "https://github.com/iamtraction/sage/releases/download/$Tag/$Artifact"

$InstallDir = "$env:LOCALAPPDATA\bin"
New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null

# download and extract binary
Write-Host "Installing sage $Version to $InstallDir"
$ZipPath = "$env:TEMP\sage.zip"
Invoke-WebRequest -Uri $Url -OutFile $ZipPath -UseBasicParsing
Expand-Archive -Path $ZipPath -DestinationPath $env:TEMP -Force
Move-Item -Path "$env:TEMP\sage.exe" -Destination "$InstallDir\sage.exe" -Force
Remove-Item $ZipPath -Force

Write-Host "Installed."
Write-Host "Ensure $InstallDir is in your PATH."
