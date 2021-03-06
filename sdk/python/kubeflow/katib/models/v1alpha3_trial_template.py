# coding: utf-8

"""
    katib

    swagger description for katib  # noqa: E501

    OpenAPI spec version: v0.1
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1alpha3TrialTemplate(object):
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
        'go_template': 'V1alpha3GoTemplate',
        'retain': 'bool'
    }

    attribute_map = {
        'go_template': 'goTemplate',
        'retain': 'retain'
    }

    def __init__(self, go_template=None, retain=None):  # noqa: E501
        """V1alpha3TrialTemplate - a model defined in Swagger"""  # noqa: E501

        self._go_template = None
        self._retain = None
        self.discriminator = None

        if go_template is not None:
            self.go_template = go_template
        if retain is not None:
            self.retain = retain

    @property
    def go_template(self):
        """Gets the go_template of this V1alpha3TrialTemplate.  # noqa: E501


        :return: The go_template of this V1alpha3TrialTemplate.  # noqa: E501
        :rtype: V1alpha3GoTemplate
        """
        return self._go_template

    @go_template.setter
    def go_template(self, go_template):
        """Sets the go_template of this V1alpha3TrialTemplate.


        :param go_template: The go_template of this V1alpha3TrialTemplate.  # noqa: E501
        :type: V1alpha3GoTemplate
        """

        self._go_template = go_template

    @property
    def retain(self):
        """Gets the retain of this V1alpha3TrialTemplate.  # noqa: E501


        :return: The retain of this V1alpha3TrialTemplate.  # noqa: E501
        :rtype: bool
        """
        return self._retain

    @retain.setter
    def retain(self, retain):
        """Sets the retain of this V1alpha3TrialTemplate.


        :param retain: The retain of this V1alpha3TrialTemplate.  # noqa: E501
        :type: bool
        """

        self._retain = retain

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
        if issubclass(V1alpha3TrialTemplate, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1alpha3TrialTemplate):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
