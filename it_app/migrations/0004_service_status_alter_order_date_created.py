# Generated by Django 4.2.7 on 2023-12-17 10:58

import datetime
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('it_app', '0003_alter_order_date_created'),
    ]

    operations = [
        migrations.AddField(
            model_name='service',
            name='status',
            field=models.IntegerField(choices=[(1, 'Действует'), (2, 'Удалена')], default=1),
        ),
        migrations.AlterField(
            model_name='order',
            name='date_created',
            field=models.DateTimeField(default=datetime.datetime(2023, 12, 17, 10, 58, 27, 467077, tzinfo=datetime.timezone.utc), verbose_name='Дата создания'),
        ),
    ]