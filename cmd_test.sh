#!/bin/bash
echo "$ go-metainspector -help"
go-metainspector --h
echo ""
echo "$ go-metainspector www.cloudcontrol.com"
go-metainspector www.cloudcontrol.com
echo ""
echo "$ go-metainspector -u www.cloudcontrol.com"
go-metainspector -u www.cloudcontrol.com
echo ""
echo "$ go-metainspector -u www.cloudcontrol.com -all"
go-metainspector -u www.cloudcontrol.com -all
echo ""
echo "$ go-metainspector fake_url"
go-metainspector fake_url
echo ""