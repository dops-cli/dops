package ping

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/sparrc/go-ping"

	"github.com/dops-cli/dops/categories"
	"github.com/dops-cli/dops/cli"
)

// Module returns the created module
type Module struct{}

// GetModuleCommands returns the commands of the module
func (Module) GetModuleCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:        "ping",
			Usage:       "Ping a host",
			Description: "Ping pings a host on the web via ICMP",
			Category:    categories.Statistics,
			Action: func(context *cli.Context) error {
				host := context.String("host")
				count := context.Int("count")

				pinger, err := ping.NewPinger(host)
				if err != nil {
					return err
				}

				pinger.SetPrivileged(true)
				pinger.Count = count
				pinger.Interval = context.Duration("interval")
				pinger.Size = context.Int("size")
				pinger.Source = context.String("source")

				// listen for ctrl-C signal
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)
				go func() {
					for range c {
						pinger.Stop()
					}
				}()

				pinger.OnRecv = func(pkt *ping.Packet) {
					fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
						pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
				}
				pinger.OnFinish = func(stats *ping.Statistics) {
					fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
					fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
						stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
					fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
						stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
				}

				fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
				pinger.Run()

				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Aliases: []string{"t", "host"},
					Name:    "target",
					Usage:   "Pings `HOST`",
				},
				&cli.IntFlag{
					Aliases: []string{"c"},
					Name:    "count",
					Usage:   "send `COUNT` pings - 0 for infinite",
					Value:   0,
				},
				&cli.DurationFlag{
					Aliases: []string{"i"},
					Name:    "interval",
					Usage:   "Sends a ping every `DURATION`",
					Value:   time.Second,
				},
				&cli.IntFlag{
					Aliases: []string{"s"},
					Name:    "size",
					Usage:   "Pings with packets of size `SIZE`",
					Value:   16,
				},
				&cli.StringFlag{
					Name:        "source",
					Usage:       "Source `IP ADDRESS` of the ping",
					DefaultText: "your hosts IP",
				},
			},
		},
	}
}
