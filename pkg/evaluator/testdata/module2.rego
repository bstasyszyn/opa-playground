package module2

import data.bindings
import data.roles

user_has_role[role_name] {
    b = bindings[_]
    b.role = role_name
    b.user = input.subject.user
}

role_has_permission[role_name] {
    r = roles[_]
    r.name = role_name
    match_with_wildcard(r.operations, input.operation)
    match_with_wildcard(r.resources, input.resource)
}

match_with_wildcard(allowed, value) {
    allowed[_] = "*"
}

match_with_wildcard(allowed, value) {
    allowed[_] = value
}
