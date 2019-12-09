package main

import (
	"os"
	"testing"
)

// For testing, following environment variable must be defined
// CHECK_MODE=TEST
// VERSION=False
// PORT=22
// USERNAME={username}
// BACKUPDEFINITION=True
// JOBNAME={jobname}
// VERBOSE=false
// COMMAND=get-job
// HOST={hostname}
// PASSWORD=[password]
// IDENTITY=[ssh private key file]

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestBEMCLI_BEJobsStatusToIcingaStatus(t *testing.T) {

	bemcli := new(BEMCLI)

	data := ` Name : tdhmmd01 - Full Weekend JobType : Backup TaskType : Full TaskName : Full Weekend IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : Saturday every 1 week(s) at 11:00 PM effective on 3/28/2019. IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=11/30/2019 11:00:02 PM; EndTime=12/1/2019 7:32:37 PM; PercentComplete=100; TotalDataSizeBytes=6722145064097; JobRateMBPerMinute=5366.038; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : tdhmmd01 - Duplicate Monthly JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Monthly IsActive : False Status : Unknown SubStatus : Ok SelectionSummary : Fully selected Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : Name : tdhmmd01 - Full Monthly JobType : Backup TaskType : Full TaskName : Full Monthly IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : First Saturday of every 1 month(s) at 11:00 PM effective on 3/28/2019. IsBackupDefinitionJob : True JobHistory : @{JobStatus=Canceled; StartTime=11/2/2019 11:00:03 PM; EndTime=11/3/2019 2:09:52 AM; PercentComplete=-1; TotalDataSizeBytes=751432335139; JobRateMBPerMinute=3890; ErrorCategory=1; ErrorCode=0; ErrorMessage=} Name : TDHMMD01-Duplicate Weekend JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Weekend IsActive : False Status : Linked SubStatus : Ok SelectionSummary : Fully selected Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=12/1/2019 7:32:47 PM; EndTime=12/2/2019 4:18:45 AM; PercentComplete=100; TotalDataSizeBytes=6722145064097; JobRateMBPerMinute=12612.89; ErrorCategory=0; ErrorCode=0; ErrorMessage=}`
	result, cond := bemcli.BEJobsStatusToIcingaStatus(data, params.verbose)
	if params.verbose {
		t.Logf("%s", result)
	}
	if cond != OK_CODE {
		t.Errorf("Error testing BEJobsStatusToIcingaStatus incorrect code %d, want 0\nResult string is %s", cond, result)
	}
	if len(bemcli.beJobStatus) != 4 {
		t.Errorf("Number of jobs incorrect %d, want 4", len(bemcli.beJobStatus))
	}

}

func TestBEMCLI_BEJobsStatusToIcingaStatus2(t *testing.T) {
	bemcli := new(BEMCLI)

	data := ` Name : TDHSAGA07-Full Weekend JobType : Backup TaskType : Full TaskName : Full Weekend IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : Friday every 1 week(s) at 11:00 PM effective on 11/24/2018. IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=11/29/2019 11:35:53 PM; EndTime=11/30/2019 2:56:26 AM; PercentComplete=100; TotalDataSizeBytes=277568567177; JobRateMBPerMinute=1447.869; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : TDHSAGA07-Full Monthly JobType : Backup TaskType : Full TaskName : Full Monthly IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : First Friday of every 1 month(s) at 11:00 PM effective on 11/24/2018. IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=11/2/2019 12:21:12 AM; EndTime=11/2/2019 4:13:04 AM; PercentComplete=100; TotalDataSizeBytes=273292245744; JobRateMBPerMinute=1282.628; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : TDHSAGA07-Duplicate Weekend JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Weekend IsActive : False Status : Linked SubStatus : Ok SelectionSummary : Fully selected Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=11/30/2019 2:56:35 AM; EndTime=11/30/2019 3:46:27 AM; PercentComplete=100; TotalDataSizeBytes=277568567177; JobRateMBPerMinute=5588.178; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : TDHSAGA07-Duplicate Monthly JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Monthly IsActive : False Status : Linked SubStatus : Ok SelectionSummary : Fully selected Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=11/2/2019 4:13:10 AM; EndTime=11/2/2019 4:53:59 AM; PercentComplete=100; TotalDataSizeBytes=273292245744; JobRateMBPerMinute=7120.065; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : TDHSAGA07-Diff Daily JobType : Backup TaskType : Differential TaskName : Diff Daily IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : Monday, Tuesday, Wednesday, Thursday every 1 week(s) at 11:00 PM effective on 4/1/2019. IsBackupDefinitionJob : True JobHistory : @{JobStatus=Succeeded; StartTime=12/3/2019 11:00:02 PM; EndTime=12/3/2019 11:08:15 PM; PercentComplete=100; TotalDataSizeBytes=4233005347; JobRateMBPerMinute=609.4657; ErrorCategory=0; ErrorCode=0; ErrorMessage=} `
	result, cond := bemcli.BEJobsStatusToIcingaStatus(data, params.verbose)
	if params.verbose {
		t.Logf("%s", result)
	}
	if cond != OK_CODE {
		t.Errorf("Error testing BEJobsStatusToIcingaStatus incorrect code %d, want 0\nResult string is %s", cond, result)
	}
	if len(bemcli.beJobStatus) != 5 {
		t.Errorf("Number of jobs incorrect %d, want 5", len(bemcli.beJobStatus))
	}
}

// Testing return code when one of jobStatus is Error
func TestBEMCLI_BEJobsStatusToIcingaStatus3(t *testing.T) {
	bemcli := new(BEMCLI)

	data := ` Name : TDHPIC01-Full Weekend JobType : Backup TaskType : Full TaskName : Full Weekend IsActive : False Status : Succeeded SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : Unscheduled IsBackupDefinitionJob : True JobHistory : @{JobStatus=Error; StartTime=10/12/2019 7:00:02 PM; EndTime=10/12/2019 9:25:15 PM; PercentComplete=0; TotalDataSizeBytes=40173961241; JobRateMBPerMinute=362.2798; ErrorCategory=7; ErrorCode=-536836884; ErrorMessage=A backup storage read/write error has occurred. If the storage is tape based, this is usually caused by dirty read/write heads in the tape drive. Clean the tape drive, and then try the job again. If the problem persists, try a different tape. You may also need to check for problems with cables, termination, or other hardware issues. If the storage is disk based, check that the storage subsystem is functioning properly. Review any system logs or vendor specific logs associated with the storage to help determine the source of the problem. You may also want to check any vendor specific documentation for troubleshooting recommendations. If the storage is cloud based, check for network connection problems. Run the CloudConnect Optimizer to obtain a value for write connections that is suitable for your environment and use this value to run the failed backup job. Review cloud provider specific documentation to help determine the source of the problem. If the problem still persists, contact the cloud provider for further assistance.} `
	result, cond := bemcli.BEJobsStatusToIcingaStatus(data, params.verbose)
	if params.verbose {
		t.Logf("%s", result)
	}
	if cond != CRI_CODE {
		t.Errorf("Error testing BEJobsStatusToIcingaStatus incorrect code %d, want 2\nResult string is %s", cond, result)
	}
	if len(bemcli.beJobStatus) != 1 {
		t.Errorf("Number of jobs incorrect %d, want 5", len(bemcli.beJobStatus))
	}
}

// Testing return code when oene jobStatus is "SucceededWithExceptions"
func TestBEMCLI_BEJobsStatusToIcingaStatus4(t *testing.T) {
	//
	bemcli := new(BEMCLI)

	data := ` Name : tdhvcenter55 Full Weekend JobType : Backup TaskType : Full TaskName : Full Weekend IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : TDHVCENTER55 Storage : Deduplication Storage 001 Schedule : Friday every 1 week(s) at 7:00 PM effective on 11/21/2018. IsBackupDefinitionJob : True JobHistory : @{JobStatus=SucceededWithExceptions; StartTime=11/29/2019 7:00:04 PM; EndTime=11/29/2019 9:59:11 PM; PercentComplete=100; TotalDataSizeBytes=86943922465; JobRateMBPerMinute=474; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : tdhvcenter55-Full Monthly JobType : Backup TaskType : Full TaskName : Full Monthly IsActive : False Status : Scheduled SubStatus : Ok SelectionSummary : TDHVCENTER55 Storage : Deduplication Storage 001 Schedule : First Friday of every 1 month(s) at 7:00 PM effective on 11/29/2018. IsBackupDefinitionJob : True JobHistory : @{JobStatus=SucceededWithExceptions; StartTime=12/6/2019 7:17:21 PM; EndTime=12/6/2019 11:07:58 PM; PercentComplete=100; TotalDataSizeBytes=87051203986; JobRateMBPerMinute=466.0001; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : tdhvcenter55-Duplicate Weekend JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Weekend IsActive : False Status : Linked SubStatus : Ok SelectionSummary : TDHVCENTER55 Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : @{JobStatus=SucceededWithExceptions; StartTime=11/29/2019 11:23:10 PM; EndTime=11/30/2019 12:02:00 AM; PercentComplete=100; TotalDataSizeBytes=86943922465; JobRateMBPerMinute=2139; ErrorCategory=0; ErrorCode=0; ErrorMessage=} Name : tdhvcenter55-Duplicate Monthly JobType : Duplicate TaskType : Duplicate TaskName : Duplicate Monthly IsActive : False Status : Linked SubStatus : Ok SelectionSummary : TDHVCENTER55 Storage : Any tape cartridge storage Schedule : IsBackupDefinitionJob : True JobHistory : @{JobStatus=SucceededWithExceptions; StartTime=12/7/2019 12:30:54 AM; EndTime=12/7/2019 1:12:20 AM; PercentComplete=100; TotalDataSizeBytes=87051203986; JobRateMBPerMinute=2068; ErrorCategory=0; ErrorCode=0; ErrorMessage=} `
	result, cond := bemcli.BEJobsStatusToIcingaStatus(data, params.verbose)
	if params.verbose {
		t.Logf("%s", result)
	}
	if cond != OK_CODE {
		t.Errorf("Error testing BEJobsStatusToIcingaStatus incorrect code %d, want 0\nResult string is %s", cond, result)
	}
	if len(bemcli.beJobStatus) != 4 {
		t.Errorf("Number of jobs incorrect %d, want 4", len(bemcli.beJobStatus))
	}

}

// Testing return code when only one Job is defined in jobDefintion and is active
func TestBEMCLI_BEJobsStatusToIcingaStatus5(t *testing.T) {
	//
	bemcli := new(BEMCLI)

	data := ` Name : TDHADM02-Full Weekend JobType : Backup TaskType : Full TaskName : Full Weekend IsActive : True Status : Active SubStatus : Ok SelectionSummary : Fully selected Storage : Deduplication Storage 001 Schedule : Friday every 1 week(s) at 11:00 PM effective on 11/24/2018. IsBackupDefinitionJob : True JobHistory : `
	result, cond := bemcli.BEJobsStatusToIcingaStatus(data, params.verbose)
	if params.verbose {
		t.Logf("%s", result)
	}
	if cond != OK_CODE {
		t.Errorf("Error testing BEJobsStatusToIcingaStatus incorrect code %d, want 0\nResult string is %s", cond, result)
	}
	if len(bemcli.beJobStatus) != 1 {
		t.Errorf("Number of jobs incorrect %d, want 1", len(bemcli.beJobStatus))
	}

}

func TestBEMCLI_Init(t *testing.T) {
	bemcli := new(BEMCLI)
	bemcli.Init(params.host, params.username, params.password, params.identity, params.port)
}

func TestBEMCLI_GetBEJobBackupDefinition(t *testing.T) {
	bemcli := new(BEMCLI)
	bemcli.Init(params.host, params.username, params.password, params.identity, params.port)
	s := bemcli.GetBEJobBackupDefinition(params.jobName)
	_ = s
}

func TestBEMCLI_GetBEBackupExecSetting(t *testing.T) {
	bemcli := new(BEMCLI)
	bemcli.Init(params.host, params.username, params.password, params.identity, params.port)
	bemcli.GetBEBackupExecSetting()
}
