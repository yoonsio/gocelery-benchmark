from worker import add_int, add_int_kwargs, add_str, add_str_kwargs

# 1) integer addition with args
add_int.apply_async((5868, 5844), serializer='json')



ar = add.apply_async((5745, 5866), serializer='json')
print('Ready status: {}'.format(ar.ready()))
print('Result: {}'.format(ar.get()))


# apply_async(args[, kwargs[, …]])
# task.apply_async(args=[arg1, arg2], kwargs={'kwarg1': 'x', 'kwarg2': 'y'})

# delay(*args, **kwargs)
# task.delay(arg1, arg2, kwarg1='x', kwarg2='y')


# calling features
# http://docs.celeryproject.org/en/latest/userguide/calling.html

# linking feature

# error handling feature


# long running task

# def on_raw_message(body):
#   print(body)
# r = long_running_task.apply_async()
# print(r.get(on_message=on_raw_message, propagate=False))

# # eta result
# # task is executed at least 3 seconds after calling
# result = add.apply_async((2, 2), countdown=3)
# result.get() # this takes at least 3 seconds to return

# tomorrow = datetime.utcnow() + timedelta(days=1)
# add.apply_async((2, 2), eta=tomorrow)

# # expiration
# add.apply_async((10, 10), expires=60)
# # worker received expired task will mark it as REVOKED


# # retry - default 3
# add.apply_async((2, 2), retry=True, retry_policy={
#     'max_retries': 3,
#     'interval_start': 0,
#     'interval_step': 0.2,
#     'interval_max': 0.2,
# })


# # json
# add.apply_async((10, 10), serializer='json')

# # compression
# add.apply_async((2, 2), compression='zlib')


# # ignore result
# result = add.apply_async(1, 2, ignore_result=True)
# result.get() # -> None

# # do not ignore result (default)
# result = add.apply_async(1, 2, ignore_result=False)
# result.get() # -> 3

# exchange - name of exchange to send message to

# set routing_key

# priority (0 - 255)