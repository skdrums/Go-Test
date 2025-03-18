package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func convertFloor(floor string) (float64, error) {
	if floor == "" {
		return 0., nil
	}
	// 正規表現で階数を判別
	re := regexp.MustCompile(`^(M)?(B)?(\d+)(F)?$`)
	matches := re.FindStringSubmatch(floor)
	fmt.Println(matches)
	if len(matches) != 5 {
		return 0, fmt.Errorf("invalid floor format: %s", floor)
	}

	// 数字部分を取得
	num, err := strconv.Atoi(matches[3])
	if err != nil {
		return 0, err
	}

	// 前置きの文字で階数を調整
	ret := float64(num)
	if matches[1] == "M" {
		// スキップフロアは0.5を減算
		ret -= 0.5
	}
	if matches[2] == "B" {
		// 地下は-1を乗算
		ret *= -1.0
	}
	return ret, nil
}

func convertFloor2(floor string) (float64, error) {
	var i, j float64 = 0, 1
	if strings.HasPrefix(floor, "M") {
		floor = floor[1:]
		i = 0.5
	}
	if strings.HasPrefix(floor, "B") {
		floor = floor[1:]
		j = -1
		i += 1 // 0階は存在しないので地下階は1加算する
	}
	v, err := strconv.ParseFloat(floor, 64)
	if err != nil {
		return 0., err
	}
	return (v - i) * j, nil
}

func main() {
	// テストケース
	floors := []string{"1F", "2F", "B1", "M2", "B3", "M3", "5F", "M5F", "MB3", "1"}

	for _, floor := range floors {
		result, err := convertFloor2(floor)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Floor %s -> %.1f\n", floor, result)
		}
	}
}
