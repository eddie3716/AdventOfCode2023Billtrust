const std = @import("std");
const q = @import("../queue.zig");
const panic = std.debug.panic;
const tokenizer = std.mem.tokenizeAny(u8, line, "->, ");

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
    name: std.string.String,
    kind: Kind,
    senders: []const std.string.String,
    next_workflows: []const std.string.String,
};

fn parseFile(file_name: []const u8) !std.HashMap(Workflow) {
    var workflows = std.HashMap(Workflow).init(std.alloc, 0);
    const contents = try std.io.readFile(file_name);
    var tokens = []const std.string.String{};
    var iter = std.mem.splitToIterator(contents, '\n');
    for (iter.next()) |line| {
        line = std.ascii.trimSuffix(line, '\n');
        tokens = std.string.split(line, |c| c == ' ' or c == ',');
        tokens = tokens.filter(s != "->" and s != "");
        const raw_name = tokens[0];
        const name = std.string.trimPrefix(raw_name, '%') | std.string.trimPrefix(raw_name, '&');
        const next_workflows = tokens[1..];

        const kind = switch (raw_name) {
            "broadcaster" => Kind.Button,
            "%" => Kind.FlipFlopper,
            "&" => Kind.Inverter,
            else => panic("Unknown kind"),
        };

        const workflow = Workflow{
            .name = name,
            .kind = kind,
            .senders = []const std.string.String{},
            .next_workflows = next_workflows,
        };
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
                    try workflow.senders.append(other_name);
                }
            }
        }
    }

    return workflows;
}

fn part1(workflows: *const std.HashMap(Workflow)) !i32 {
    var lowPulses: i32 = 0;
    var highPulses: i32 = 0;
    var queue = q.Queue(QueueItem).init(std.heap.page_allocator);

    const onOrOff = std.HashMap(bool).init(std.alloc, 0);
    const inverterSignals = std.HashMap(Signal).init(std.alloc, 0);

    for (workflows.entries) |workflow| {
        if (workflow.kind == Kind.FlipFlopper) {
            try onOrOff.put(workflow.name, false);
        }
    }

    const button = "button";

    for (1..1000) |i| {
        lowPulses += 1;

        var visited = std.HashMap(bool).init(std.alloc, 0);

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
            var newSignal: Signal;

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
            } else if (currentWorkflowName == "broadcaster") {
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
