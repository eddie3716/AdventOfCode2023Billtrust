const std = @import("std");
const q = @import("../queue.zig");
const b = @import("../buf-iter.zig");
const panic = std.debug.panic;

const Signal = enum {
    High,
    Low,
};

const Kind = enum {
    Button,
    FlipFlopper,
    Inverter,
};

const QueueItem = struct {
    signal: Signal,
    from_name: *const std.string.String,
    to_workflow: *const Workflow,
};

const Workflow = struct {
    name: []const u8,
    kind: Kind,
    senders: std.ArrayList([]const u8),
    next_workflows: std.ArrayList([]const u8),
};

fn parseFile(file_name: []const u8) !*std.StringHashMap(Workflow) {
    var iterator = try b.iterLines(file_name);
    defer iterator.deinit();
    var workflows = &std.StringHashMap(Workflow).init(std.heap.page_allocator);
    while (try iterator.next()) |line| {
        var tokenizer = std.mem.tokenizeAny(u8, line, "->, ");

        const raw_name = tokenizer.next().?;
        const name = if (raw_name[0] == '%' or raw_name[0] == '&') raw_name[1..] else raw_name;
        const kind = switch (raw_name[0]) {
            '%' => Kind.FlipFlopper,
            '&' => Kind.Inverter,
            else => Kind.Button,
        };

        const workflow = Workflow{
            .name = name,
            .kind = kind,
            .senders = std.ArrayList([]const u8).init(std.heap.page_allocator),
            .next_workflows = std.ArrayList([]const u8).init(std.heap.page_allocator),
        };

        while (try tokenizer.next()) |token| {
            workflow.next_workflows.append(token);
        }

        const insert = workflows.put(name, workflow);
        if (insert == null) {
            panic("Failed to insert workflow into HashMap");
        }
    }

    for (workflows) |workflow| {
        const name = workflow.name;
        if (workflow.kind == Kind.Inverter) {
            for (workflows) |other_workflow| {
                const other_workflow_name = other_workflow.name;
                if (std.array.contains(other_workflow.next_workflows, name)) {
                    try workflow.senders.append(other_workflow_name);
                }
            }
        }
    }

    return workflows;
}

fn part1(workflows: *const std.StringHashMap(Workflow)) !i32 {
    var lowPulses: i32 = 0;
    var highPulses: i32 = 0;
    var queue = q.Queue(QueueItem).init(std.heap.page_allocator);

    const onOrOff = std.StringHashMap(bool).init(std.heap.page_allocator);
    const inverterSignals = std.StringHashMap(Signal).init(std.heap.page_allocator);

    for (workflows.entries) |workflow| {
        if (workflow.kind == Kind.FlipFlopper) {
            try onOrOff.put(workflow.name, false);
        }
    }

    const button = "button";

    for (1..1000) |i| {
        _ = i;
        lowPulses += 1;

        var visited = std.StringHashMap(bool).init(std.heap.page_allocator);

        const broadcaster = try workflows.get("broadcaster");
        if (broadcaster != null) {
            const queueItem = QueueItem{
                .signal = Signal.Low,
                .from_name = button,
                .to_workflow = broadcaster,
            };
            try queue.enqueue(queueItem);
        } else {
            panic("No broadcaster found");
        }

        while (queue.dequeue()) |queueItem| {
            const toWorkflow = queueItem.to_workflow;
            const fromName = queueItem.from_name;
            const currentSignal = queueItem.signal;
            const currentKind = &toWorkflow.kind;
            const currentWorkflowName = toWorkflow.name;
            var newSignal: Signal = undefined;

            if (currentKind == Kind.FlipFlopper) {
                const entry = onOrOff.mutate(currentWorkflowName);
                if (currentSignal == Signal.Low) {
                    entry.* = !entry.*;
                }
                if (entry.*) {
                    newSignal = Signal.High;
                } else {
                    newSignal = Signal.Low;
                }
            } else if (currentKind == Kind.Inverter) {
                const inverterKey = currentWorkflowName ++ "_" ++ fromName;
                const inverterSignal = inverterSignals.get(inverterKey);
                if (inverterSignal == Signal.Low) {
                    newSignal = Signal.Low;
                } else {
                    newSignal = Signal.High;
                }

                if (currentSignal == Signal.Low) {
                    inverterSignals.put(inverterKey, Signal.Low);
                } else {
                    inverterSignals.put(inverterKey, Signal.High);
                }
            } else if (std.mem.eql(u8, currentWorkflowName, "broadcaster")) {
                newSignal = Signal.Low;
            } else {
                if (currentSignal == Signal.Low) {
                    newSignal = Signal.Low;
                } else {
                    newSignal = Signal.High;
                }
            }

            for (toWorkflow.next_workflows) |nextWorkflowName| {
                const newSignalStr = switch (newSignal) {
                    Signal.Low => "low",
                    Signal.High => "high",
                };

                const visitedKey = currentWorkflowName ++ "_" ++ nextWorkflowName ++ "_" ++ newSignalStr;
                if (visited.get(visitedKey)) {
                    continue;
                } else {
                    visited.put(visitedKey, true);
                }

                if (newSignal == Signal.Low) {
                    lowPulses += 1;
                } else {
                    highPulses += 1;
                }

                const nextWorkflow = workflows.get(nextWorkflowName);
                if (nextWorkflow != null) {
                    const newQueueItem = QueueItem{
                        .signal = newSignal,
                        .from_name = &toWorkflow.name,
                        .to_workflow = nextWorkflow,
                    };
                    try queue.append(newQueueItem);
                }
            }
        }
    }

    std.debug.print("low_pulses: {}\n", .{lowPulses});
    std.debug.print("high_pulses: {}\n", .{highPulses});
    return lowPulses * highPulses;
}

fn part2(workflows: *const std.HashMap(Workflow)) !i32 {
    _ = workflows;
    var sum: i32 = 0;
    // Implement part 2 logic here if needed.
    return sum;
}

pub fn main() void {
    const file_name = "testinput.txt";
    const workflows = try parseFile(file_name);
    const result1 = try part1(&workflows);
    const result2 = try part2(&workflows);

    std.debug.print("part1: {}\n", .{result1});
    std.debug.print("part2: {}\n", .{result2});
}
