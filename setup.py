# Copyright 2020 Belobragin V.V.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from setuptools import setup

# Do not edit these constants. They will be updated automatically
# by scripts/update-client.sh.
CLIENT_VERSION = "0.0.4"
PACKAGE_NAME = "bvv_utils"
DEVELOPMENT_STATUS = "4 - Beta"

# To install the library, run the following
#
# python setup.py install
#
# prerequisite: setuptools
# http://pypi.python.org/pypi/setuptools

REQUIRES = []
with open('requirements.txt') as f:
    for line in f:
        line, _, _ = line.partition('#')
        line = line.strip()
        if ';' in line:
            requirement, _, specifier = line.partition(';')
            for_specifier = EXTRAS.setdefault(':{}'.format(specifier), [])
            for_specifier.append(requirement)
        else:
            REQUIRES.append(line)

setup(
    name=PACKAGE_NAME,
    version=CLIENT_VERSION,
    description="Useful utils for Kaggle image processing competitions",
    author_email="vblvbl133@gmail.com",
    author="Belobragin Vladimir",
    license="Apache License Version 2.0",
    url="https://github.com/Belobragin/bvv_utils",
    keywords=["BVV"],
    install_requires=REQUIRES,
    #packages=['bvv_utils'],
    py_modules=['bvv_utils'],
    include_package_data=True,
    long_description="Python utilites for Kaggle image processing competitions, no further description",
    classifiers=[
        "Development Status :: %s" % DEVELOPMENT_STATUS,
        "Topic :: Utilities",
        "Intended Audience :: BV private",
        "Intended Audience :: Kaggle BV projects",
        "License :: OSI Approved :: Apache Software License",
        "Operating System :: Centos7",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8 - ?",
    ],
)
