package module1

import data.module2.user_has_role
import data.module2.role_has_permission

import data.bindings
import data.roles

default allow = false

allow {
    custom_print("******* Executing custom builtin function *******")
    user_has_role[role_name]
    role_has_permission[role_name]
}

not_allow {
    not allow
}
