    if latest.ko.Status.VPCLinkStatus != nil && *latest.ko.Status.VPCLinkStatus != svcsdk.VpcLinkStatusAvailable {
        return nil, waitForAvailableRequeue
    }