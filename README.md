# Social Food (API)

This API is used social-food

## Run
```markdown
$ docker-compose -f docker-compose-db.yml up {{-d}}   
$ docker-compose up {{--force-recreate}} {{-d}}
```

### Option description
> -f: 기본 docker-compose 파일 말고 실행하기 위함  
> -d: 백그라운드에서 실행  
> --force-recreate : 새 이미지 생성해서 실행

## Git Hooks Setup
### Required
```shell
$ cp -R git-hooks/* .git/hooks
$ chmod ug+x .git/hooks/*
```

### Optional
If you want to colorized console, execute below.
```shell
$ sudo gem install colorize 
$ sudo gem install win32console  # For Windows only
```
If not, delete colorized sentence/word in `git-hooks`