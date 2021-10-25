    // Ignore deploymentId when autodeploy is set to true
    if *input.AutoDeploy == true {
        input.DeploymentId = nil
    }