# Licensing Overview

Scheduling Service is offered under a dual-license model so you can choose the option that best fits how you deploy the software.

## License Texts

- [GNU Affero General Public License v3.0](../LICENSE)
- [Scheduling Service Commercial License Terms](commercial-license.md)

## Obligations Under the GNU AGPL v3

Using the community edition under the GNU Affero General Public License v3.0 means:

- **Source availability:** If you distribute the software or make it available over a network (including as a hosted service) with modifications, you must provide the complete corresponding source code of those modifications to your users.
- **Copyleft scope:** Any derivative work that you distribute must also be licensed under the AGPL v3 so downstream users inherit the same freedoms.
- **No warranty:** The software is provided “as-is” with no warranty, as described in the license text.

## Commercial Licensing Options

A commercial license allows you to use Scheduling Service without the copyleft obligations of the AGPL. Typical benefits include:

- Keeping proprietary modifications private while distributing or hosting the service.
- Receiving commercial support, additional warranties, and negotiated service levels.
- Combining Scheduling Service with third-party proprietary components without triggering reciprocal obligations.

Contact the licensing team at `sales@example.com` to request pricing or receive a tailored agreement.

## Self-Hosting for Free Without Source Changes

You may deploy and operate Scheduling Service internally without charge under the AGPL as long as you do **not** modify the source code. Running the unmodified software for yourself does not create any obligation to publish your internal configurations or workflows. If you later modify the source, review the AGPL requirements above to ensure you provide those changes to users who can access the modified service.

## Typical Usage Scenarios

| Scenario | What it looks like | AGPL obligations | When to consider commercial licensing |
| --- | --- | --- | --- |
| **Open-source deployment** | Running the unmodified community edition or contributing improvements back to the project. | You can self-host freely. If you distribute builds or host modified code, share those modifications under the AGPL. | Usually not required unless you want to keep changes proprietary. |
| **Internal fork** | Customizing the code for internal teams or subsidiaries only. | If employees or subsidiaries access a modified hosted service, provide them with the corresponding source under the AGPL. | Needed when you prefer to keep the fork private or integrate with closed-source components. |
| **SaaS offering** | Providing Scheduling Service as part of a hosted product to paying customers. | Customers must receive access to the modified source if you change it. Network use triggers the AGPL’s sharing requirements. | Recommended to avoid disclosing proprietary enhancements or mixing with non-AGPL-compatible code. |

## Additional Guidance

- Document your deployment architecture and keep track of any code changes to simplify compliance.
- If you distribute client libraries or SDKs alongside the service, ensure their licenses are compatible with your chosen model.
- Consult legal counsel for definitive advice before launching commercial offerings.
