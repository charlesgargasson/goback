from setuptools import setup, find_packages
setup(
    name='goback',
    version='1.1.7',
    packages=find_packages(),
    install_requires=[
        'docker',
        'netifaces',
        'gitpython',
        'colorama'
    ],
    entry_points={
        'console_scripts': [
            'goback=pycli.goback:main',
        ],
    },
)