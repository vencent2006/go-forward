import random
from datetime import datetime, timedelta

# 定义可能的客户姓名、商品名称、渠道
customer_names = ["张三", "李四", "王五", "赵六", "孙七", "周八", "吴九", "郑十"]
product_names = ["手机", "电脑", "耳机", "键盘", "鼠标", "显示器", "打印机", "摄像头"]
channels = ["微信", "支付宝", "拼多多", "小红书"]
titles = ["订单编号", "客户姓名", "商品名称", "商品数量", "单价", "收入", "渠道", "订单日期"]

# 生成 2025 年 2 月的日期列表
start_date = datetime(2025, 2, 1)
end_date = datetime(2025, 2, 28)
date_list = [start_date + timedelta(days=x) for x in range((end_date - start_date).days + 1)]
# 随机选取 100 个日期
random_dates = [random.choice(date_list) for _ in range(100)]

# 对选取的日期进行升序排列
sorted_dates = sorted(random_dates)

# 将日期转换为字符串格式
sorted_dates_str = [date.strftime("%Y-%m-%d") for date in sorted_dates]

# 生成 100 行订单数据
orders = []
for i in range(1, 101):
    order_id = i
    customer_name = random.choice(customer_names)
    product_name = random.choice(product_names)
    quantity = random.randint(1, 5)
    price = random.randint(100, 10000)
    income = quantity * price
    order_date = sorted_dates_str[i - 1]
    channel = random.choice(channels)
    orders.append([order_id, customer_name, product_name, quantity, price, income, channel, order_date])

for title in titles:
    print(title, end=" ")
print()
for order in orders:
    for item in order:
        print(item, end=" ")
    print()

exit(0)
# 写入 CSV 文件
with open(
        'data.csv',
        'w', newline='', encoding='utf-8') as csvfile:
    writer = csv.writer(csvfile)
    # 写入表头
    writer.writerow(["订单编号", "客户姓名", "商品名称", "商品数量", "单价", "收入", "渠道", "订单日期"])
    # 写入订单数据
    writer.writerows(orders)
