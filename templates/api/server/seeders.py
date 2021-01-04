from datetime import datetime

from bolinette import blnt
from bolinette.decorators import seeder
from bolinette.defaults.services import RoleService, UserService

@seeder
async def role_seeder(context: blnt.BolinetteContext):
    role_service: RoleService = context.service('role')
    async with blnt.Transaction(context):
        await role_service.create({'name': 'root'})
        await role_service.create({'name': 'admin'})


@seeder
async def dev_user_seeder(context: blnt.BolinetteContext):
    if context.env['profile'] == 'development':
        role_service: RoleService = context.service('role')
        user_service: UserService = context.service('user')
        async with blnt.Transaction(context):
            root = await role_service.get_by_name('root')
            admin = await role_service.get_by_name('admin')
            root_usr = await user_service.create({
                'username': 'root',
                'password': 'root',
                'email': f'root@localhost'
            })
            root_usr.roles.append(root)
            root_usr.roles.append(admin)

            dev0 = await role_service.create({'name': 'dev0'})
            dev1 = await role_service.create({'name': 'dev1'})
            dev2 = await role_service.create({'name': 'dev2'})
            roles = [dev0, dev1, dev2]

            for i in range(10):
                user = await user_service.create({
                    'username': f'user_{i}',
                    'password': 'test',
                    'email': f'user{i}@test.com'
                })
                user.roles.append(roles[i % 3])
                user.roles.append(roles[(i + 1) % 3])
