const std = @import("std");
const buf = @import("../buf-iter.zig");
const q = @import("../queue.zig");
const io = std.io;
const fs = std.fs;
const ArrayList = std.ArrayList;
var gpa = std.heap.GeneralPurposeAllocator(.{}){};
const BROADCASTER = "broadcaster";

pub fn parseFile(fileName: []const u8) !StateMachine {
    var iterator = try buf.iterLines(fileName);
    defer iterator.deinit();

    var stateMachine = StateMachine{
        .states = std.StringHashMap(States).init(gpa.allocator()),
        .workflows = std.StringHashMap(Workflow).init(gpa.allocator()),
    };

    while (try iterator.next()) |line| {
        var tokenizer = std.mem.tokenizeAny(u8, line, "->, ");

        var workflow = try initWorkflow(gpa.allocator());

        var firstToken = tokenizer.next().?;

        workflow.workflowType = if (std.mem.eql(u8, firstToken, BROADCASTER)) WorkflowTypes.Button else if (firstToken[0] == '%') WorkflowTypes.FlipFlopper else WorkflowTypes.Inverter;
        while (tokenizer.next()) |token| {
            try workflow.nextWorkflows.append(token);
        }

        if (std.mem.eql(u8, firstToken, BROADCASTER)) {
            //std.debug.print("{s} to {s}\n", .{ firstToken, workflow.nextWorkflows.items });
            workflow.name = firstToken;
            try stateMachine.states.put(firstToken, States.Low);
            try stateMachine.workflows.put(firstToken, workflow);
        } else {
            //std.debug.print("Setting {s} to low\n", .{firstToken[1..]});
            workflow.name = firstToken[1..];
            try stateMachine.states.put(firstToken[1..], States.Low);
            try stateMachine.workflows.put(firstToken[1..], workflow);
        }

        std.debug.print("{s}\n", .{line});
    }

    return stateMachine;
}

const Unknown = 0;
const Button = 1;
const FlipFlopper = 2;
const Inverter = 3;

const States = enum { High, Low };
const WorkflowTypes = enum { Unknown, Button, FlipFlopper, Inverter };

const Workflow = struct {
    name: []const u8,
    workflowType: WorkflowTypes,
    nextWorkflows: ArrayList([]const u8),
};

const StateMachine = struct {
    states: std.StringHashMap(States),
    workflows: std.StringHashMap(Workflow),
};

pub fn initWorkflow(allocator: std.mem.Allocator) !Workflow {
    return Workflow{
        .name = undefined,
        .workflowType = WorkflowTypes.Unknown,
        .nextWorkflows = ArrayList([]const u8).init(allocator),
    };
}

pub fn main() !void {
    var stateMachine = try parseFile("./testinput.txt");

    std.debug.print("Part 1: {any}\n", .{part1(stateMachine)});
    std.debug.print("Part 2: {any}\n", .{part2()});
}

pub fn part1(statemachine: StateMachine) !i32 {
    var answer: i32 = 0;

    var queue = q.Queue(Workflow).init(gpa.allocator());

    try queue.enqueue(statemachine.workflows.get(BROADCASTER).?);

    var workflowPtr = queue.dequeue();

    while (workflowPtr != null) {
        var workflow = workflowPtr.?;
        switch (workflow.workflowType) {
            WorkflowTypes.Button => {
                for (workflow.nextWorkflows.items) |nextWorkflow| {
                    var state = statemachine.states.get(nextWorkflow).?;
                    if (state == States.High) {
                        try statemachine.states.put(nextWorkflow, States.Low);
                    } else {
                        try statemachine.states.put(nextWorkflow, States.High);
                    }
                    try queue.enqueue(statemachine.workflows.get(nextWorkflow).?);
                    answer += 1;
                }
            },
            WorkflowTypes.FlipFlopper => {
                for (workflow.nextWorkflows.items) |nextWorkflow| {
                    var state = statemachine.states.get(nextWorkflow).?;
                    if (state == States.High) {
                        try statemachine.states.put(nextWorkflow, States.Low);
                    } else {
                        try statemachine.states.put(nextWorkflow, States.High);
                    }
                    try queue.enqueue(statemachine.workflows.get(nextWorkflow).?);
                    answer += 1;
                }
            },
            WorkflowTypes.Inverter => {
                for (workflow.nextWorkflows.items) |nextWorkflow| {
                    var state = statemachine.states.get(nextWorkflow).?;
                    if (state == States.High) {
                        try statemachine.states.put(nextWorkflow, States.Low);
                    } else {
                        try statemachine.states.put(nextWorkflow, States.High);
                    }
                    try queue.enqueue(statemachine.workflows.get(nextWorkflow).?);
                    answer += 1;
                }
            },
            else => unreachable,
        }
        workflowPtr = queue.dequeue();
    }

    return answer;
}

pub fn part2() !i32 {
    var answer: i32 = 0;

    return answer;
}
