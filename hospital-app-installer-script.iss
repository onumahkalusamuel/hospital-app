; -- Example1.iss --
; Demonstrates copying 3 files and creating an icon.

; SEE THE DOCUMENTATION FOR DETAILS ON CREATING .ISS SCRIPT FILES!

[Setup]
AppName=Hospital Card Mgmt. System
AppVersion=1.1.0
WizardStyle=modern
DefaultDirName={autopf}\HCMS
DefaultGroupName=HCMS
UninstallDisplayIcon={app}\hcms.exe
Compression=lzma2
SolidCompression=yes
OutputDir=userdocs:HCMS

[Files]
Source: "hcms.exe"; DestDir: "{app}"
Source: "applauncher.exe"; DestDir: "{app}"
Source: "hcms.ico"; DestDir: "{app}"

[Icons]
Name: "{userdesktop}\Hospital Card Mgmt. System"; Filename: "{app}\applauncher.exe"; IconFilename: "{app}\hcms.ico"
Name: "{group}\Hospital Card Mgmt. System"; Filename: "{app}\applauncher.exe"; IconFilename: "{app}\hcms.ico"

[Run]
Filename: {sys}\sc.exe; Parameters: "create HcmsSvc start= auto type= share binPath= ""{app}\hcms.exe"" DisplayName= ""HCMS"""; Flags: runhidden
Filename: "{sys}\sc.exe"; Parameters: "start HcmsSvc"; Flags: runhidden
Filename: "{sys}\netsh.exe"; Parameters: "firewall set portopening protocol=TCP port=""{code:GetServerPort}"" name=NxTCP mode=ENABLE"; StatusMsg: "Opening TCP Port ""{code:GetServerPort}"; Flags: runhidden

[Code]
function GetServerPort(Value: string): String;
begin
  Result := '8788';
end;

[UninstallRun]
Filename: {sys}\sc.exe; Parameters: "stop HcmsSvc" ; Flags: runhidden; RunOnceId: "stopService"
Filename: {sys}\sc.exe; Parameters: "delete HcmsSvc" ; Flags: runhidden; RunOnceId: "delService"
