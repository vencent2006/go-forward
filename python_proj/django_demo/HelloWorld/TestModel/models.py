from django.db import models


# Create your models here.
class entLocalWhitelist(models.Model):
    # 模型在数据库中的表名默认是 APP名称_类名，可以通过db_table重新默认值，相当于指定模型在数据库中的表明
    # db_table = 'ent_local_whitelist'
    id = models.IntegerField(primary_key=True)
    uid = models.BigIntegerField()
    class Meta:
        # 模型在数据库中的表名默认是 APP名称_类名，可以通过db_table重新默认值，相当于指定模型在数据库中的表明
        db_table = 'ent_local_whitelist'