resource "aws_ecs_cluster" "main" {
  name = "go-pf"
}

resource "aws_security_group" "alb" {
  name        = "go-pf-alb"
  description = "go-pf alb"
  vpc_id      = (aws_vpc.main.id)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "go-pf-alb"
  }
}

resource "aws_security_group_rule" "alb_http" {
  security_group_id = (aws_security_group.alb.id)
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
}

resource "aws_lb" "main" {
  load_balancer_type = "application"
  name               = "go-pf"
  security_groups = [(aws_security_group.alb.id)]
  subnets         = [(aws_subnet.public_1a.id), (aws_subnet.public_1c.id), (aws_subnet.public_1d.id)]
}

resource "aws_lb_listener" "main" {
  port     = "80"
  protocol = "HTTP"
  load_balancer_arn = (aws_lb.main.arn)
  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      status_code  = "200"
      message_body = "ok"
    }
  }
}

resource "aws_lb_target_group" "main" {
  name   = "go-pf"
  vpc_id = (aws_vpc.main.id)
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  health_check {
    port = 80
    path = "/"
  }
}

resource "aws_lb_listener_rule" "main" {
  listener_arn = (aws_lb_listener.main.arn)
  action {
    type             = "forward"
    target_group_arn = (aws_lb_target_group.main.id)
  }
  condition {
    path_pattern {
      values = ["*"]
    }
  }
}
resource "aws_internet_gateway" "main" {
  vpc_id = (aws_vpc.main.id)
  tags = {
    Name = "go-pf"
  }
}

resource "aws_route_table" "public" {
  vpc_id = (aws_vpc.main.id)
  tags = {
    Name = "go-pf-public"
  }
}

resource "aws_route" "public" {
  destination_cidr_block = "0.0.0.0/0"
  route_table_id         = (aws_route_table.public.id)
  gateway_id             = (aws_internet_gateway.main.id)
}

resource "aws_route_table_association" "public_1a" {
  subnet_id      = (aws_subnet.public_1a.id)
  route_table_id = (aws_route_table.public.id)
}

resource "aws_route_table_association" "public_1c" {
  subnet_id      = (aws_subnet.public_1c.id)
  route_table_id = (aws_route_table.public.id)
}

resource "aws_route_table_association" "public_1d" {
  subnet_id      = (aws_subnet.public_1d.id)
  route_table_id = (aws_route_table.public.id)
}

resource "aws_route_table" "private_1a" {
  vpc_id = (aws_vpc.main.id)
  tags = {
    Name = "go-pf-private-1a"
  }
}

resource "aws_route_table" "private_1c" {
  vpc_id = (aws_vpc.main.id)
  tags = {
    Name = "go-pf-private-1c"
  }
}

resource "aws_route_table" "private_1d" {
  vpc_id = (aws_vpc.main.id)
  tags = {
    Name = "go-pf-private-1d"
  }
}

resource "aws_security_group" "ecs" {
  name        = "go-pf-ecs"
  description = "go-pf ecs"
  vpc_id      = (aws_vpc.main.id)
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "go-pf-ecs"
  }
}

resource "aws_security_group_rule" "ecs" {
  security_group_id = (aws_security_group.ecs.id)
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["10.0.0.0/16"]
}

resource "aws_ecs_service" "main" {
  name       = "go-pf"
  depends_on = [aws_lb_listener_rule.main]
  cluster = (aws_ecs_cluster.main.name)
  launch_type     = "FARGATE"
  desired_count   = "1"
  task_definition = (aws_ecs_task_definition.main.arn)
  network_configuration {
    subnets         = [(aws_subnet.private_1a.id), (aws_subnet.private_1c.id), (aws_subnet.private_1d.id)]
    security_groups = [(aws_security_group.ecs.id)]
  }
  load_balancer {
    target_group_arn = (aws_lb_target_group.main.arn)
    container_name   = "nginx"
    container_port   = "80"
  }
}

resource "aws_subnet" "public_1a" {
  vpc_id            = (aws_vpc.main.id)
  availability_zone = "ap-northeast-1a"
  cidr_block        = "10.0.1.0/24"
  tags = {
    Name = "go-pf-public-1a"
  }
}

resource "aws_subnet" "private_1a" {
  vpc_id            = (aws_vpc.main.id)
  availability_zone = "ap-northeast-1a"
  cidr_block        = "10.0.10.0/24"
  tags = {
    Name = "go-pf-private-1a"
  }
}

resource "aws_subnet" "public_1c" {
  vpc_id = (aws_vpc.main.id)

  availability_zone = "ap-northeast-1c"

  cidr_block = "10.0.2.0/24"

  tags = {
    Name = "go-pf-public-1c"
  }
}

resource "aws_subnet" "private_1c" {
  vpc_id = (aws_vpc.main.id)

  availability_zone = "ap-northeast-1c"
  cidr_block        = "10.0.20.0/24"

  tags = {
    Name = "go-pf-private-1c"
  }
}

resource "aws_subnet" "public_1d" {
  vpc_id = (aws_vpc.main.id)

  availability_zone = "ap-northeast-1d"

  cidr_block = "10.0.3.0/24"

  tags = {
    Name = "go-pf-public-1d"
  }
}

resource "aws_subnet" "private_1d" {
  vpc_id = (aws_vpc.main.id)

  availability_zone = "ap-northeast-1d"
  cidr_block        = "10.0.30.0/24"

  tags = {
    Name = "go-pf-private-1d"
  }
}

resource "aws_ecs_task_definition" "main" {
  family = "go-pf"
  requires_compatibilities = ["FARGATE"]
  cpu    = "256"
  memory = "512"
  network_mode = "awsvpc"
  container_definitions = <<EOL
[
  {
    "name": "nginx",
    "image": "nginx:1.14",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 80
      }
    ]
  }
]
EOL
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "go-pf"
  }
}