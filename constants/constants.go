package constants

import (
	"go.uber.org/zap"
	"os"
)

type Constantstruct struct {
	EnvVar map[string]string
}

var Constants = Constantstruct{
	EnvVar: map[string]string{
		"BACKLOG_TITLE_PREFIX":   os.Getenv("BACKLOG_TITLE_PREFIX"),
		"BACKLOG_BODY_FILE_PATH": os.Getenv("BACKLOG_BODY_FILE_PATH"),
		"AGED_TITLE_PREFIX":      os.Getenv("AGED_TITLE_PREFIX"),
		"AGED_BODY_FILE_PATH":    os.Getenv("AGED_BODY_FILE_PATH"),
		"FDR_TITLE_PREFIX":       os.Getenv("FDR_TITLE_PREFIX"),
		"FDR_BODY_FILE_PATH":     os.Getenv("FDR_BODY_FILE_PATH"),
	},
}

func CheckAndSetEnvVars() {
	// Set default values if the environment variables are not set
	if Constants.EnvVar["BACKLOG_TITLE_PREFIX"] == "" {
		Constants.EnvVar["BACKLOG_TITLE_PREFIX"] = "ansalemo | backlog review"
		zap.L().Warn("BACKLOG_TITLE_PREFIX is not set, using default value  of " + "'" + Constants.EnvVar["BACKLOG_TITLE_PREFIX"] + "'")
	} else {
		zap.L().Info("BACKLOG_TITLE_PREFIX is set to " + "'" + Constants.EnvVar["BACKLOG_TITLE_PREFIX"] + "'")
	}

	if Constants.EnvVar["AGED_TITLE_PREFIX"] == "" {
		Constants.EnvVar["AGED_TITLE_PREFIX"] = "ansalemo | aged case review"
		zap.L().Warn("AGED_TITLE_PREFIX is not set, using default value of " + "'" + Constants.EnvVar["AGED_TITLE_PREFIX"] + "'")
	} else {
		zap.L().Info("AGED_TITLE_PREFIX is set to " + "'" + Constants.EnvVar["AGED_TITLE_PREFIX"] + "'")
	}

	if Constants.EnvVar["FDR_TITLE_PREFIX"] == "" {
		Constants.EnvVar["FDR_TITLE_PREFIX"] = "ansalemo | FDR review"
		zap.L().Warn("FDR_TITLE_PREFIX is not set, using default value of " + "'" + Constants.EnvVar["FDR_TITLE_PREFIX"] + "'")
	} else {
		zap.L().Info("FDR_TITLE_PREFIX is set to " + "'" + Constants.EnvVar["FDR_TITLE_PREFIX"] + "'")
	}

	// If this environment variable isn't set then lookup from the current directory
	if Constants.EnvVar["BACKLOG_BODY_FILE_PATH"] == "" {
		Constants.EnvVar["BACKLOG_BODY_FILE_PATH"] = "./backlog_body.txt"
		zap.L().Warn("BACKLOG_BODY_FILE_PATH is not set, looking up from the current directory of " + "'" + Constants.EnvVar["BACKLOG_BODY_FILE_PATH"] + "'")
	} else {
		zap.L().Info("BACKLOG_BODY_FILE_PATH is set to " + "'" + Constants.EnvVar["BACKLOG_BODY_FILE_PATH"] + "'")
	}

	// If this environment variable isn't set then lookup from the current directory
	if Constants.EnvVar["AGED_BODY_FILE_PATH"] == "" {
		Constants.EnvVar["AGED_BODY_FILE_PATH"] = "./aged_body.txt"
		zap.L().Warn("AGED_BODY_FILE_PATH is not set, looking up from the current directory of " + "'" + Constants.EnvVar["AGED_BODY_FILE_PATH"] + "'")
	} else {
		zap.L().Info("AGED_BODY_FILE_PATH is set to " + "'" + Constants.EnvVar["AGED_BODY_FILE_PATH"] + "'")
	}

	// If this environment variable isn't set then lookup from the current directory
	if Constants.EnvVar["FDR_BODY_FILE_PATH"] == "" {
		Constants.EnvVar["FDR_BODY_FILE_PATH"] = "./fdr_body.txt"
		zap.L().Warn("FDR_BODY_FILE_PATH is not set, looking up from the current directory of " + "'" + Constants.EnvVar["FDR_BODY_FILE_PATH"] + "'")
	} else {
		zap.L().Info("FDR_BODY_FILE_PATH is set to " + "'" + Constants.EnvVar["FDR_BODY_FILE_PATH"] + "'")
	}
}
