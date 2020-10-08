package ping

import (
	"os"
	"os/signal"
	"time"

	"github.com/go-ping/ping"
	"github.com/pterm/pterm"

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

				pb := pterm.DefaultProgressbar.WithTotal(count).WithTitle("Ping").Start()

				// listen for ctrl-C signal
				c := make(chan os.Signal, 1)
				signal.Notify(c, os.Interrupt)
				go func() {
					for range c {
						pinger.Stop()
					}
				}()

				pinger.OnRecv = func(pkt *ping.Packet) {
					pterm.Printf("%s bytes from %s: icmp_seq: %s time: %v\n", pterm.LightMagenta(pkt.Nbytes), pterm.Yellow(pkt.IPAddr), pterm.LightMagenta(pkt.Seq), pterm.LightMagenta(pkt.Rtt))
					pb.Increment()
				}
				pinger.OnFinish = func(stats *ping.Statistics) {
					pterm.Printf("\n" + pterm.Gray("       --- ") + pterm.LightWhite(stats.Addr+" ping statistics ") + pterm.Gray("---") + "\n")
					pterm.Printf("%s packets transmitted, %s packets received, %v packet loss\n", pterm.LightMagenta(stats.PacketsSent), pterm.LightMagenta(stats.PacketsRecv), pterm.LightMagenta(stats.PacketLoss, "%"))
					pterm.Printf("round-trip "+pterm.Green("min")+"/"+pterm.Yellow("avg")+"/"+pterm.Red("max")+"/"+pterm.LightMagenta("stddev")+" = %v/%v/%v/%v\n", pterm.Green(stats.MinRtt), pterm.Yellow(stats.AvgRtt), pterm.Red(stats.MaxRtt), pterm.LightMagenta(stats.StdDevRtt))
				}

				pterm.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
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
