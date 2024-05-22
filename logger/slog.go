/*
 * @Author: i@douxuefeng.cn
 * @Date: 2024-05-21 15:30:17
 * @LastEditTime: 2024-05-21 15:31:06
 * @LastEditors: i@douxuefeng.cn
 * @Description:
 */
package logger

import (
	"log/slog"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitLogger(dir string) {
	dest := path.Join(dir, time.Now().Format(time.DateOnly)+".log")
	f, _ := rotatelogs.New(
		dest,
		rotatelogs.WithLinkName(dir+"logs.log"),
		// rotatelogs.WithMaxAge(time.Duration(viper.GetInt("logMaxAge")*24)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	slog.SetDefault(slog.New(slog.NewJSONHandler(f, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	})))
}
