import os

GOOD_AND_GOARCHS = [
    ("darwin", "amd64",),
    ("linux", "amd64",),
    ("windows", "amd64",)
]


for i in GOOD_AND_GOARCHS:
    platform, arch = i
    folder = f"delfin0-0-1_{platform}-{arch}"
    os.system(f"set GOOS={platform}&&set GOARCH={arch}&&mkdir {folder}&&go build delfin.go handler.go structs.go&&move {'./delfin' if platform != 'windows' else 'delfin.exe'} {folder}&&move {folder} build")