# Generated by Django 4.2.7 on 2023-12-18 19:15

import datetime
from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('it_app', '0007_alter_order_date_created'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='order',
            name='id_user',
        ),
        migrations.AlterField(
            model_name='order',
            name='date_created',
            field=models.DateTimeField(default=datetime.datetime(2023, 12, 18, 19, 15, 53, 966934, tzinfo=datetime.timezone.utc), verbose_name='Дата создания'),
        ),
    ]