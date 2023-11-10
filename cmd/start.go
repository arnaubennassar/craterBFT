package main

import (
	"github.com/urfave/cli/v2"
)

func StartNode(cli *cli.Context) error {

	// 	baseContext := context.Background()

	// 	homeDir := processPath(cli.String("home"))
	// 	// jwtFile := processPath(cli.String("jwt"))
	// 	// if _, err := os.Stat(jwtFile); err != nil {
	// 	// 	return err
	// 	// }

	// 	// engineUrl := cli.String("engineUrl")
	// 	// ethUrl := cli.String("ethUrl")

	// 	cfg := config.DefaultConfig()
	// 	cfg.SetRoot(homeDir)

	// 	var err error
	// 	viper.SetConfigFile(fmt.Sprintf("%s/%s", homeDir, "config/config.toml"))
	// 	if err = viper.ReadInConfig(); err != nil {
	// 		return err
	// 	}
	// 	if err = viper.Unmarshal(cfg); err != nil {
	// 		return err
	// 	}
	// 	if err = cfg.ValidateBasic(); err != nil {
	// 		return err
	// 	}

	// 	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	// 	if logger, err = flags.ParseLogLevel(cfg.LogLevel, logger, config.DefaultLogLevel); err != nil {
	// 		return err
	// 	}

	// 	// state := app.NewState(filepath.Join(homeDir, "data"))
	// 	// defer state.Close()

	// 	// logger.Info("eth API", "url", ethUrl)
	// 	// logger.Info("engine API", "url", engineUrl)
	// 	// logger.Info("engine JWT", "file", jwtFile)

	// 	pv := privval.LoadFilePV(
	// 		cfg.PrivValidatorKeyFile(),
	// 		cfg.PrivValidatorStateFile(),
	// 	)

	// 	addr := common.BytesToAddress(pv.GetAddress().Bytes())

	// 	var nodeKey *p2p.NodeKey
	// 	if nodeKey, err = p2p.LoadNodeKey(cfg.NodeKeyFile()); err != nil {
	// 		return err
	// 	}

	// 	eth, err := ethclient.DialContext(baseContext, ethUrl)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	eng, err := engine.NewEngineClient(engineUrl, jwtFile)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	service := app.NewService(eth, eng, logger, cfg.Moniker, addr, state)
	// 	err = service.CheckCapabilities()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	var n *node.Node
	// 	if n, err = node.NewNode(
	// 		cfg,
	// 		pv,
	// 		nodeKey,
	// 		proxy.NewLocalClientCreator(service),
	// 		node.DefaultGenesisDocProviderFunc(cfg),
	// 		config.DefaultDBProvider,
	// 		node.DefaultMetricsProvider(cfg.Instrumentation),
	// 		logger); err != nil {
	// 		return err
	// 	}

	// 	if err = n.Start(); err != nil {
	// 		return err
	// 	}

	// 	defer func() {
	// 		_ = n.Stop()
	// 		n.Wait()
	// 	}()

	// 	c := make(chan os.Signal, 1)
	// 	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// 	<-c

	// 	return err
	// }

	//	func processPath(path string) string {
	//		if userHome, err := os.UserHomeDir(); err == nil {
	//			path = strings.Replace(path, "~", userHome, 1)
	//			path = filepath.Clean(path)
	//		}
	//		return path
	return nil
}
