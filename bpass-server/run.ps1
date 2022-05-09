

param (
    [String]$Pack ,
    [String]$Type = "run"
)
$name="bpass-server"
$curDir= Split-Path -Parent $MyInvocation.MyCommand.Definition
Write-Host "当前路径"$curDir   -ForegroundColor Yellow
if ($Type -eq "run") {
    Write-Host "运行" -ForegroundColor Cyan
    if ($Os -eq "win") {
        gf run  main.go -p bin
    }
    else {
        Write-Host "系统不匹配" -ForegroundColor Red
    }

}
elseif ($Type -eq "build") {
    Write-Host "编译" -ForegroundColor Red
    

        gf build main.go


}