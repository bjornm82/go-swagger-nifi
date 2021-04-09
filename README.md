# Swagger NiFi

## Example usage

### Create Process Group (on root)

[POST] http://localhost:8080/nifi-api/process-groups/b2cc3e2c-0178-1000-375e-8b7c9361724c/process-groups

request
```
{"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":0},"disconnectedNodeAcknowledged":false,"component":{"name":"datasource-1","position":{"x":207,"y":242}}}
```

[201] response
```
{"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":1,"lastModifier":"anonymous"},"id":"b6701c5b-0178-1000-0000-0000261eb470","uri":"http://localhost:8080/nifi-api/process-groups/b6701c5b-0178-1000-0000-0000261eb470","position":{"x":207.0,"y":242.0},"permissions":{"canRead":true,"canWrite":true},"bulletins":[],"component":{"id":"b6701c5b-0178-1000-0000-0000261eb470","parentGroupId":"b2cc3e2c-0178-1000-375e-8b7c9361724c","position":{"x":207.0,"y":242.0},"name":"datasource-1","comments":"","variables":{},"flowfileConcurrency":"UNBOUNDED","flowfileOutboundPolicy":"STREAM_WHEN_AVAILABLE","runningCount":0,"stoppedCount":0,"invalidCount":0,"disabledCount":0,"activeRemotePortCount":0,"inactiveRemotePortCount":0,"upToDateCount":0,"locallyModifiedCount":0,"staleCount":0,"locallyModifiedAndStaleCount":0,"syncFailureCount":0,"localInputPortCount":0,"localOutputPortCount":0,"publicInputPortCount":0,"publicOutputPortCount":0,"contents":{"processGroups":[],"remoteProcessGroups":[],"processors":[],"inputPorts":[],"outputPorts":[],"connections":[],"labels":[],"funnels":[],"controllerServices":[]},"inputPortCount":0,"outputPortCount":0},"status":{"id":"b6701c5b-0178-1000-0000-0000261eb470","name":"datasource-1","statsLastRefreshed":"11:41:43 UTC","aggregateSnapshot":{"id":"b6701c5b-0178-1000-0000-0000261eb470","name":"datasource-1","flowFilesIn":0,"bytesIn":0,"input":"0 (0 bytes)","flowFilesQueued":0,"bytesQueued":0,"queued":"0 (0 bytes)","queuedCount":"0","queuedSize":"0 bytes","bytesRead":0,"read":"0 bytes","bytesWritten":0,"written":"0 bytes","flowFilesOut":0,"bytesOut":0,"output":"0 (0 bytes)","flowFilesTransferred":0,"bytesTransferred":0,"transferred":"0 (0 bytes)","bytesReceived":0,"flowFilesReceived":0,"received":"0 (0 bytes)","bytesSent":0,"flowFilesSent":0,"sent":"0 (0 bytes)","activeThreadCount":0,"terminatedThreadCount":0}},"runningCount":0,"stoppedCount":0,"invalidCount":0,"disabledCount":0,"activeRemotePortCount":0,"inactiveRemotePortCount":0,"upToDateCount":0,"locallyModifiedCount":0,"staleCount":0,"locallyModifiedAndStaleCount":0,"syncFailureCount":0,"localInputPortCount":0,"localOutputPortCount":0,"publicInputPortCount":0,"publicOutputPortCount":0,"inputPortCount":0,"outputPortCount":0}
```

### Create AWSCredentialsProviderControllerService

[POST] http://localhost:8080/nifi-api/process-groups/b6701c5b-0178-1000-0000-0000261eb470/controller-services

request
```
{"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":0},"disconnectedNodeAcknowledged":false,"component":{"type":"org.apache.nifi.processors.aws.credentials.provider.service.AWSCredentialsProviderControllerService","bundle":{"group":"org.apache.nifi","artifact":"nifi-aws-nar","version":"1.13.0"}}}
```

[201] response
```
{"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":1,"lastModifier":"anonymous"},"id":"b67181b3-0178-1000-ffff-ffffdf6972c5","uri":"http://localhost:8080/nifi-api/controller-services/b67181b3-0178-1000-ffff-ffffdf6972c5","permissions":{"canRead":true,"canWrite":true},"bulletins":[],"parentGroupId":"b6701c5b-0178-1000-0000-0000261eb470","component":{"id":"b67181b3-0178-1000-ffff-ffffdf6972c5","parentGroupId":"b6701c5b-0178-1000-0000-0000261eb470","name":"AWSCredentialsProviderControllerService","type":"org.apache.nifi.processors.aws.credentials.provider.service.AWSCredentialsProviderControllerService","bundle":{"group":"org.apache.nifi","artifact":"nifi-aws-nar","version":"1.13.0"},"controllerServiceApis":[{"type":"org.apache.nifi.processors.aws.credentials.provider.service.AWSCredentialsProviderService","bundle":{"group":"org.apache.nifi","artifact":"nifi-aws-service-api-nar","version":"1.13.0"}}],"state":"DISABLED","persistsState":false,"restricted":false,"deprecated":false,"multipleVersionsAvailable":false,"properties":{"default-credentials":null,"Access Key":null,"Secret Key":null,"Credentials File":null,"profile-name":null,"anonymous-credentials":null,"Assume Role ARN":null,"Assume Role Session Name":null,"Session Time":null,"assume-role-external-id":null,"assume-role-proxy-host":null,"assume-role-proxy-port":null,"assume-role-sts-endpoint":null},"descriptors":{"default-credentials":{"name":"default-credentials","displayName":"Use Default Credentials","description":"If true, uses the Default Credential chain, including EC2 instance profiles or roles, environment variables, default user credentials, etc.","defaultValue":"false","allowableValues":[{"allowableValue":{"displayName":"true","value":"true"},"canRead":true},{"allowableValue":{"displayName":"false","value":"false"},"canRead":true}],"required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Access Key":{"name":"Access Key","displayName":"Access Key ID","description":"","required":false,"sensitive":true,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"Secret Key":{"name":"Secret Key","displayName":"Secret Access Key","description":"","required":false,"sensitive":true,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"Credentials File":{"name":"Credentials File","displayName":"Credentials File","description":"Path to a file containing AWS access key and secret key in properties file format.","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"profile-name":{"name":"profile-name","displayName":"Profile Name","description":"The AWS profile name for credentials from the profile configuration file.","required":false,"sensitive":false,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"anonymous-credentials":{"name":"anonymous-credentials","displayName":"Use Anonymous Credentials","description":"If true, uses Anonymous credentials","defaultValue":"false","allowableValues":[{"allowableValue":{"displayName":"true","value":"true"},"canRead":true},{"allowableValue":{"displayName":"false","value":"false"},"canRead":true}],"required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Assume Role ARN":{"name":"Assume Role ARN","displayName":"Assume Role ARN","description":"The AWS Role ARN for cross account access. This is used in conjunction with role name and session timeout","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Assume Role Session Name":{"name":"Assume Role Session Name","displayName":"Assume Role Session Name","description":"The AWS Role Name for cross account access. This is used in conjunction with role ARN and session time out","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Session Time":{"name":"Session Time","displayName":"Session Time","description":"Session time for role based session (between 900 and 3600 seconds). This is used in conjunction with role ARN and name","defaultValue":"3600","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-external-id":{"name":"assume-role-external-id","displayName":"Assume Role External ID","description":"External ID for cross-account access. This is used in conjunction with role arn, role name, and optional session time out","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-proxy-host":{"name":"assume-role-proxy-host","displayName":"Assume Role Proxy Host","description":"Proxy host for cross-account access, if needed within your environment. This will configure a proxy to request for temporary access keys into another AWS account","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-proxy-port":{"name":"assume-role-proxy-port","displayName":"Assume Role Proxy Port","description":"Proxy port for cross-account access, if needed within your environment. This will configure a proxy to request for temporary access keys into another AWS account","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-sts-endpoint":{"name":"assume-role-sts-endpoint","displayName":"Assume Role STS Endpoint","description":"The default AWS Security Token Service (STS) endpoint (\"sts.amazonaws.com\") works for all accounts that are not for China (Beijing) region or GovCloud. You only need to set this property to \"sts.cn-north-1.amazonaws.com.cn\" when you are requesting session credentials for services in China(Beijing) region or to \"sts.us-gov-west-1.amazonaws.com\" for GovCloud.","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]}},"validationErrors":[],"validationStatus":"VALID","extensionMissing":false},"operatePermissions":{"canRead":true,"canWrite":true},"status":{"runStatus":"DISABLED","validationStatus":"VALID"}}
```

Update Controller Service

[PUT] http://localhost:8080/nifi-api/controller-services/b67181b3-0178-1000-ffff-ffffdf6972c5

request
```
{"disconnectedNodeAcknowledged":false,"component":{"id":"b67181b3-0178-1000-ffff-ffffdf6972c5","name":"AWSCredentialsProviderControllerService","comments":"","properties":{"Access Key":"minio","Secret Key":"minio123"}},"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":1}}
```

[200] response
```
{"revision":{"clientId":"b65ef9b7-0178-1000-be42-5e7e6b9c7d9a","version":2,"lastModifier":"anonymous"},"id":"b67181b3-0178-1000-ffff-ffffdf6972c5","uri":"http://localhost:8080/nifi-api/controller-services/b67181b3-0178-1000-ffff-ffffdf6972c5","permissions":{"canRead":true,"canWrite":true},"bulletins":[],"parentGroupId":"b6701c5b-0178-1000-0000-0000261eb470","component":{"id":"b67181b3-0178-1000-ffff-ffffdf6972c5","parentGroupId":"b6701c5b-0178-1000-0000-0000261eb470","name":"AWSCredentialsProviderControllerService","type":"org.apache.nifi.processors.aws.credentials.provider.service.AWSCredentialsProviderControllerService","bundle":{"group":"org.apache.nifi","artifact":"nifi-aws-nar","version":"1.13.0"},"controllerServiceApis":[{"type":"org.apache.nifi.processors.aws.credentials.provider.service.AWSCredentialsProviderService","bundle":{"group":"org.apache.nifi","artifact":"nifi-aws-service-api-nar","version":"1.13.0"}}],"comments":"","state":"DISABLED","persistsState":false,"restricted":false,"deprecated":false,"multipleVersionsAvailable":false,"properties":{"default-credentials":null,"Access Key":"********","Secret Key":"********","Credentials File":null,"profile-name":null,"anonymous-credentials":null,"Assume Role ARN":null,"Assume Role Session Name":null,"Session Time":null,"assume-role-external-id":null,"assume-role-proxy-host":null,"assume-role-proxy-port":null,"assume-role-sts-endpoint":null},"descriptors":{"default-credentials":{"name":"default-credentials","displayName":"Use Default Credentials","description":"If true, uses the Default Credential chain, including EC2 instance profiles or roles, environment variables, default user credentials, etc.","defaultValue":"false","allowableValues":[{"allowableValue":{"displayName":"true","value":"true"},"canRead":true},{"allowableValue":{"displayName":"false","value":"false"},"canRead":true}],"required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Access Key":{"name":"Access Key","displayName":"Access Key ID","description":"","required":false,"sensitive":true,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"Secret Key":{"name":"Secret Key","displayName":"Secret Access Key","description":"","required":false,"sensitive":true,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"Credentials File":{"name":"Credentials File","displayName":"Credentials File","description":"Path to a file containing AWS access key and secret key in properties file format.","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"profile-name":{"name":"profile-name","displayName":"Profile Name","description":"The AWS profile name for credentials from the profile configuration file.","required":false,"sensitive":false,"dynamic":false,"supportsEl":true,"expressionLanguageScope":"Variable Registry Only","dependencies":[]},"anonymous-credentials":{"name":"anonymous-credentials","displayName":"Use Anonymous Credentials","description":"If true, uses Anonymous credentials","defaultValue":"false","allowableValues":[{"allowableValue":{"displayName":"true","value":"true"},"canRead":true},{"allowableValue":{"displayName":"false","value":"false"},"canRead":true}],"required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Assume Role ARN":{"name":"Assume Role ARN","displayName":"Assume Role ARN","description":"The AWS Role ARN for cross account access. This is used in conjunction with role name and session timeout","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Assume Role Session Name":{"name":"Assume Role Session Name","displayName":"Assume Role Session Name","description":"The AWS Role Name for cross account access. This is used in conjunction with role ARN and session time out","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"Session Time":{"name":"Session Time","displayName":"Session Time","description":"Session time for role based session (between 900 and 3600 seconds). This is used in conjunction with role ARN and name","defaultValue":"3600","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-external-id":{"name":"assume-role-external-id","displayName":"Assume Role External ID","description":"External ID for cross-account access. This is used in conjunction with role arn, role name, and optional session time out","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-proxy-host":{"name":"assume-role-proxy-host","displayName":"Assume Role Proxy Host","description":"Proxy host for cross-account access, if needed within your environment. This will configure a proxy to request for temporary access keys into another AWS account","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-proxy-port":{"name":"assume-role-proxy-port","displayName":"Assume Role Proxy Port","description":"Proxy port for cross-account access, if needed within your environment. This will configure a proxy to request for temporary access keys into another AWS account","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]},"assume-role-sts-endpoint":{"name":"assume-role-sts-endpoint","displayName":"Assume Role STS Endpoint","description":"The default AWS Security Token Service (STS) endpoint (\"sts.amazonaws.com\") works for all accounts that are not for China (Beijing) region or GovCloud. You only need to set this property to \"sts.cn-north-1.amazonaws.com.cn\" when you are requesting session credentials for services in China(Beijing) region or to \"sts.us-gov-west-1.amazonaws.com\" for GovCloud.","required":false,"sensitive":false,"dynamic":false,"supportsEl":false,"expressionLanguageScope":"Not Supported","dependencies":[]}},"referencingComponents":[],"validationErrors":[],"validationStatus":"VALID","extensionMissing":false},"operatePermissions":{"canRead":true,"canWrite":true},"status":{"runStatus":"DISABLED","validationStatus":"VALID"}}
```


## Generate Go Client

Download the release of nifi `https://github.com/apache/nifi/releases`

unzip / untar the file

`cd /{nifi-release}/nifi-nar-bundles/nifi-framework-bundle/nifi-framework/nifi-web/nifi-web-api`

mvn clean package

`cd target/swagger-ui`

you should find the `swagger.json` file

mv the file to the root of this project

`swagger generate client`

Most likely some errors occur when generating:

```
- "getState" is defined 4 times
- "updateRunStatus" is defined 5 times
- "clearState" is defined 3 times
- "createControllerService" is defined 2 times
- "getPropertyDescriptor" is defined 3 times
- "removeDropRequest" is defined 2 times
- "deleteUpdateRequest" is defined 2 times
- operation "uploadTemplate" has both formData and body parameters. Only one such In: type may be used for a given operation
```
Resolving manually by renaming the operationId's by prefixing the tags camelCased.

e.g.
```
"tags" : [ "controller-services" ],
"summary" : "Gets the state for a controller service",
"description" : "",
"operationId" : "getState", <-- orginal
```
should become:
```
"tags" : [ "controller-services" ],
"summary" : "Gets the state for a controller service",
"description" : "",
"operationId" : "controllerServicesGetState", <-- to become
```

Add a miniumum to RevisionDTO.Version, this will create an *int64, so 0 can be set (default unmarshalling behaviour golang)
```
"RevisionDTO" : {
    "type" : "object",
    "properties" : {
    "version" : {
        "format" : "int64",
        "minimum": 0
```

Bug in the code of NiFi, when posting a 201 status code is returned, not a 200, so all the posts needs to have a 201.

TODO: UploadTemplate for now I removed the body, which isn't tested at all, should followup on this