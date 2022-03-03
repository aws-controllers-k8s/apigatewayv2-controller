    // if the Target was provided using a reference, prepend "integrations/"
    // with the IntegrationId from TargetRef
    if ko.Spec.TargetRef != nil && ko.Spec.Target != nil {
        targetStr := fmt.Sprintf("integrations/%s", *ko.Spec.Target)
        ko.Spec.Target = &targetStr
    }
