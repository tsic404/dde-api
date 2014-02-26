/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package main

import (
        "dlib/dbus"
)

type Graph struct {
        BlurPictChanged func(string, string)
}

func (graph *Graph) GetDBusInfo() dbus.DBusInfo {
        return dbus.DBusInfo{
                "com.deepin.api.Graph",
                "/com/deepin/api/Graph",
                "com.deepin.api.Graph",
        }
}

func main() {
        defer func() {
                if err := recover(); err != nil {
                        // TODO logFatal("deepin graph api failed: %v", err)
                }
        }()

        jobInHand = make(map[string]bool) // used by blur pict

        graph := &Graph{}
        err := dbus.InstallOnSession(graph)
        if err != nil {
                panic(err)
        }
        dbus.DealWithUnhandledMessage()

        select {}
}
