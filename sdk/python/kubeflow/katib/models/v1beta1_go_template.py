# coding: utf-8

"""
    Katib

    Swagger description for Katib  # noqa: E501

    OpenAPI spec version: v1beta1-0.1
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1beta1GoTemplate(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'raw_template': 'str',
        'template_spec': 'V1beta1TemplateSpec'
    }

    attribute_map = {
        'raw_template': 'rawTemplate',
        'template_spec': 'templateSpec'
    }

    def __init__(self, raw_template=None, template_spec=None):  # noqa: E501
        """V1beta1GoTemplate - a model defined in Swagger"""  # noqa: E501

        self._raw_template = None
        self._template_spec = None
        self.discriminator = None

        if raw_template is not None:
            self.raw_template = raw_template
        if template_spec is not None:
            self.template_spec = template_spec

    @property
    def raw_template(self):
        """Gets the raw_template of this V1beta1GoTemplate.  # noqa: E501


        :return: The raw_template of this V1beta1GoTemplate.  # noqa: E501
        :rtype: str
        """
        return self._raw_template

    @raw_template.setter
    def raw_template(self, raw_template):
        """Sets the raw_template of this V1beta1GoTemplate.


        :param raw_template: The raw_template of this V1beta1GoTemplate.  # noqa: E501
        :type: str
        """

        self._raw_template = raw_template

    @property
    def template_spec(self):
        """Gets the template_spec of this V1beta1GoTemplate.  # noqa: E501


        :return: The template_spec of this V1beta1GoTemplate.  # noqa: E501
        :rtype: V1beta1TemplateSpec
        """
        return self._template_spec

    @template_spec.setter
    def template_spec(self, template_spec):
        """Sets the template_spec of this V1beta1GoTemplate.


        :param template_spec: The template_spec of this V1beta1GoTemplate.  # noqa: E501
        :type: V1beta1TemplateSpec
        """

        self._template_spec = template_spec

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1GoTemplate):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
