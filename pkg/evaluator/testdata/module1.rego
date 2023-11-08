#
#   Copyright Gen Digital Inc. All Rights Reserved.
#
#   SPDX-License-Identifier: Apache-2.0
#

package module1

import data.module2.user_has_role
import data.module2.role_has_permission

import data.bindings
import data.roles

default allow = false

allow {
    user_has_role[role_name]
    role_has_permission[role_name]
}

not_allow {
    not allow
}

custom_func {
    custom_print("******* Executing custom builtin function *******")

    true
}
