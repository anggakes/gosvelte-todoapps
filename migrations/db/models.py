
from django.contrib.postgres.fields import JSONField
from django.db import models

class Todo(models.Model):
    title = models.CharField(max_length=128)
    description = models.CharField(max_length=128)
    created_at = models.DateTimeField()
    updated_at = models.DateTimeField()

    class Meta:
        db_table = '"todos"'