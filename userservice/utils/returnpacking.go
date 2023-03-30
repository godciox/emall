package utils

import pb "userservice/proto"

func PackInfo(err error, status string, desc string, response *pb.UserResponse) {
	response.Status = status
	response.Description = desc
}
