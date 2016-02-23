OPENCONTROL_V2_SCHEMA = {
  "type": "map",
  "mapping": {
    "name": {
      "type": "str",
      "required": True
    },
    "key": {
      "type": "str"
    },
    "schema_version": {
      "type": "number",
      "required": True
    },
    "documentation_complete": {
      "type": "bool"
    },
    "system": {
      "type": "str",
      "required": False
    },
    "references": {
      "type": "seq",
      "sequence": [
        {
          "type": "map",
          "mapping": {
            "name": {
              "type": "str",
              "required": True
            },
            "path": {
              "type": "str"
            },
            "type": {
              "type": "str",
              "required": True
            }
          }
        }
      ]
    },
    "verifications": {
      "type": "seq",
      "sequence": [
        {
          "type": "map",
          "mapping": {
            "key": {
              "type": "str",
              "required": True
            },
            "name": {
              "type": "str",
              "required": True
            },
            "path": {
              "type": "str"
            },
            "type": {
              "type": "str",
              "required": True
            },
            "steps": {
              "type": "str"
            },
            "filename": {
              "type": "str"
            },
            "test_passed": {
              "type": "bool"
            },
            "last_run": {
              "type": "any"
            },
          }
        }
      ]
    },
    "satisfies": {
      "type": "seq",
      "sequence": [
        {
          "type": "map",
          "mapping": {
            "standard_key": {
              "type": "text",
              "required": True
            },
            "control_key": {
              "type": "text",
              "required": True
            },
            "narrative": {
              "type": "any",
              "required": True
            },
            "implementation_status": {
              "type": "str",
              "enum": [
                "partial",
                "complete",
                "planned",
                "none"
              ]
            },
            "covered_by": {
              "type": "seq",
              "sequence": [
                {
                  "type": "map",
                  "mapping": {
                    "system_key": {
                      "type": "str"
                    },
                    "component_key": {
                      "type": "str"
                    },
                    "verification_key": {
                      "type": "str",
                      "required": True
                    }
                  }
                }
              ]
            }
          }
        }
      ]
    }
  }
}
