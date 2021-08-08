$repo = "JSONhilder/strongbox"
#Users directory by default
$default_path = "$home\.strongbox" 

$releases = "https://api.github.com/repos/$repo/releases"

Write-Host Determining latest release

$tag = (Invoke-WebRequest $releases | ConvertFrom-Json)[0].tag_name

$file = "strongbox_" + $tag.replace("v", "") +"_windows_amd64.zip"

$download = "https://github.com/$repo/releases/download/$tag/$file"

$name = $file.replace(".zip", "")

$zip = $file

Write-Host Dowloading latest release
Invoke-WebRequest $download -Out $zip

Write-Host Creating directory at $default_path
New-Item -ItemType Directory -Force -Path $default_path

Write-Host Extracting release files
Expand-Archive $zip -Force -DestinationPath $default_path

# Cleaning up target dir
Remove-Item $name -Recurse -Force -ErrorAction SilentlyContinue 

# Removing temp files
Remove-Item $zip -Force

$exe_path = $default_path+"\strongbox.exe"

$alias_exists = Test-Path alias:strongbox
if($alias_exists -eq $False) {
    Write-Host Creating alias to strongbox executable
    $alias_path = "`nSet-Alias strongbox " + $exe_path
    Add-Content -Path $profile -Value $alias_path
    & $profile
}
