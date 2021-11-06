# EVSE states

## Observed states and possible descriptions

| State | Possible description |
|-|-
| 0 | Wallbox starting |
| 1 ] Ready, no vehicle connected |
| 2 | Not ready, but vehicle connected |
| 3 | Unknown |
| 4 | Vehicle connected but not charging |
| 5 | Unknown |
| 6 | Unknown |
| 7 | Unknown |
| 8 | Vehicle connected, fully charged |
| 9 | Vehicle connected, waitig to charge |
| 10 | Unknown |
| 11 | Vehicle connected, charging with full power |
| 12 | Unknown |

## Observed patterns

### Scheduled charge

When a charge is scheduled the EVSE state is 9. After the vehicle goes to sleep it's set to 4. When charging starts the state jumps to 11.

### Charge start

When connecting the vehicle and it starts to charge the pattern is 2 -> 4 -> 12 -> 11
