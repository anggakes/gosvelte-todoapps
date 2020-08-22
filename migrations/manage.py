import sys
import os

from django.conf import settings
from django.core.management import execute_from_command_line

settings.configure(
    DEBUG=True,
    INSTALLED_APPS=[
        'db',
    ],
    DATABASES={
        'default': {
            'ENGINE': 'django.db.backends.postgresql_psycopg2',
            'NAME': os.getenv('POSTGRES_DB', 'membership'),
            'USER': os.getenv('POSTGRES_USER', 'admin'),
            'PASSWORD': os.getenv('POSTGRES_PASSWORD', 'dev-password'),
            'HOST': os.getenv('POSTGRES_HOST', 'localhost'),
            'PORT': os.getenv('POSTGRES_PORT', '5433'),
        }
    },
)

if __name__ == '__main__':
    execute_from_command_line(sys.argv)
