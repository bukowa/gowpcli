package fixorientation

//Fix image orientation for one or more attachments.

type FixOrientation struct {
    
    AttachmentId []string // [<attachment-id>...]
    
    DryRun bool // [--dry-run]
    
}

//## OPTIONS
//
//[<attachment-id>...]
//: One or more IDs of the attachments to regenerate.
//
//[--dry-run]
//: Check images needing orientation without performing the operation.
//
//## EXAMPLES
//
//    # Fix orientation for all images.
//    $ wp media fix-orientation
//    1/3 Fixing orientation for "Landscape_4" (ID 62).
//    2/3 Fixing orientation for "Landscape_3" (ID 61).
//    3/3 Fixing orientation for "Landscape_2" (ID 60).
//    Success: Fixed 3 of 3 images.
//
//    # Fix orientation dry run.
//    $ wp media fix-orientation 63 -dry run
//    1/1 "Portrait_6" (ID 63) will be affected.
//    Success: 1 of 1 image will be affected.
//
//    # Fix orientation for specific images.
//    $ wp media fix-orientation 63
//    1/1 Fixing orientation for "Portrait_6" (ID 63).
//    Success: Fixed 1 of 1 images.
//
//