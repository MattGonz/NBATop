# NBATop
A [gocui](https://github.com/jroimartin/gocui) terminal UI for NBA stats!

### Build from source
```
go get github.com/MattGonz/NBATop
cd $GOPATH/src/github.com/MattGonz/NBATop
go build -o nbatop main.go
./nbatop
```


#### Todo
###### Views
- [x] Games today
- [x] Standings
- [x] Team game logs
- [x] Game box scores
- [x] Player game logs
- [ ] Player profiles
- [ ] Multiple seasons
- [ ] `?` → help view


###### UI improvements
- [ ] Games today → nba.com matchup pages (view?)
- [ ] Horizontal scrolling
- [ ] UI formatting keybinds
- [ ] Table sorting
- [ ] Fuzzing finding



##### Inspirations
* [jesseduffield/lazygit](https://github.com/jesseduffield/lazygit) and [lazydocker](https://github.com/jesseduffield/lazydocker)
* [ClementTsang/bottom](https://github.com/ClementTsang/bottom)
* [miguelmota/cointop](https://github.com/cointop-sh/cointop)
* The [Unofficial NBA API documentation](http://nbasense.com/nba-api/) of [jasonroman/nba-api](https://github.com/jasonroman/nba-api)
